package main

import (
	"UTS/controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/roomDetail", controller.GetAllRoomDetail).Methods("GET")
	router.HandleFunc("/insertRoom", controller.GetAllRoomDetail).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
