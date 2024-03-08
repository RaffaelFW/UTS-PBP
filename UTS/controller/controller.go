package controller

import (
	m "UTS/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	query := "SELECT * FROM Rooms"
	roomName := r.URL.Query()["room_name"]
	gameID := r.URL.Query()["game_id"]

	if roomName != nil {
		fmt.Println(roomName[0])
		query += " WHERE room_name= '" + roomName[0] + "'"
	}

	if gameID != nil {
		if roomName[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " game_id= '" + gameID[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var room m.Room
	var rooms []m.Room
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName, &room.GameID); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	var response m.RoomsResponse
	response.Status = http.StatusOK
	response.Message = "Rooms retrieved successfully"
	response.Data = rooms

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
func GetAllRoomDetail(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	query := "SELECT * FROM Rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error querying rooms:", err)
		http.Error(w, "Failed to retrieve room details", http.StatusInternalServerError)
		return
	}

	var room m.Room
	var rooms []m.Room
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName, &room.GameID); err != nil {
			log.Println("Error scanning room row:", err)
			http.Error(w, "Failed to retrieve room details", http.StatusInternalServerError)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	var roomDetailResponses []m.RoomDetailResponse
	for _, room := range rooms {
		participants, err := getRoomParticipants(room.ID)
		if err != nil {
			log.Println("Error retrieving room participants:", err)
			http.Error(w, "Failed to retrieve room details", http.StatusInternalServerError)
			return
		}

		roomDetailResp := m.RoomDetailResp{
			Status: http.StatusOK,
			Data: m.RoomDetailData{
				Room:         room,
				Participants: participants,
			},
		}

		var roomDetailResponse m.RoomDetailResponse
		roomDetailResponse.Status = roomDetailResp.Status
		roomDetailResponse.Data.Room.ID = roomDetailResp.Data.Room.ID
		roomDetailResponse.Data.Room.RoomName = roomDetailResp.Data.Room.RoomName
		roomDetailResponse.Data.Room.Participants = roomDetailResp.Data.Participants
		roomDetailResponses = append(roomDetailResponses, roomDetailResponse)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roomDetailResponses)
}

func getRoomParticipants(roomID int) ([]m.Participant, error) {
	db := connectDB()
	defer db.Close()

	query := "SELECT * FROM participants WHERE room_id = ?"

	rows, err := db.Query(query, roomID)
	if err != nil {
		return nil, err
	}

	var participant m.Participant
	var participants []m.Participant
	for rows.Next() {
		if err := rows.Scan(&participant.ID, &participant.RoomID, &participant.AccountID); err != nil {
			return nil, err
		} else {
			participants = append(participants, participant)
		}
	}

	return participants, nil
}

func InsertRoomParticipant(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var roomParticipant m.Participant
	err := json.NewDecoder(r.Body).Decode(&roomParticipant)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	participants, err := getRoomParticipants(roomParticipant.RoomID)
	if err != nil {
		log.Println("Error retrieving room participants:", err)
		http.Error(w, "Failed to check room participants", http.StatusInternalServerError)
		return
	}

	game, err := getGameByRoomID(roomParticipant.RoomID)
	if err != nil {
		log.Println("Error retrieving game information:", err)
		http.Error(w, "Failed to retrieve game information", http.StatusInternalServerError)
		return
	}

	if len(participants) >= game.MaxPlayers {
		response := m.ParticipantResponse{
			Status:  http.StatusBadRequest,
			Message: "Room is full. Cannot add more participants.",
			Data:    m.Participant{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	query := "INSERT INTO participants (room_id, account_id) VALUES (?, ?)"
	_, err = db.Exec(query, roomParticipant.RoomID, roomParticipant.AccountID)
	if err != nil {
		log.Println("Error inserting participant into room:", err)
		http.Error(w, "Failed to insert participant into room", http.StatusInternalServerError)
		return
	}

	response := m.ParticipantResponse{
		Status:  http.StatusOK,
		Message: "Participant added to the room successfully.",
		Data:    roomParticipant,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func getGameByRoomID(roomID int) (m.Game, error) {
	db := connectDB()
	defer db.Close()

	query := "SELECT Games.id, Games.name, Games.max_players FROM Games INNER JOIN Rooms ON Games.id = Rooms.game_id WHERE Rooms.id = ?"

	var game m.Game
	err := db.QueryRow(query, roomID).Scan(&game.ID, &game.Name, &game.MaxPlayers)
	if err != nil {
		return m.Game{}, err
	}

	return game, nil
}
