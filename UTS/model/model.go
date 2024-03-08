package model

// Account model
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// AccountResponse model untuk respons Account
type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

// AccountsResponse model untuk respons daftar Account
type AccountsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

// Game model
type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"max_players"`
}

// Room model
type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	GameID   int    `json:"game_id"`
}
type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}
type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type Participant struct {
	ID        int `json:"id"`
	RoomID    int `json:"room_id"`
	AccountID int `json:"account_id"`
}

// GameResponse model untuk respons Game
type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

// GamesResponse model untuk respons daftar Game
type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

// ParticipantResponse model untuk respons Participant
type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type RoomDetailData struct {
	Room         Room          `json:"room"`
	Participants []Participant `json:"participants"`
}
type RoomDetailResp struct {
	Status int            `json:"status"`
	Data   RoomDetailData `json:"data"`
}

type RoomDetailResponse struct {
	Status int `json:"status"`
	Data   struct {
		Room struct {
			ID           int           `json:"id"`
			RoomName     string        `json:"room_name"`
			Participants []Participant `json:"participants"`
		} `json:"room"`
	} `json:"data"`
}
