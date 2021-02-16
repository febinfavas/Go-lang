package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

var db *gorm.DB
var err error

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	db.Create(&booking)
	fmt.Println("Endpoint Hit: Creating New Booking")
	json.NewEncoder(w).Encode(booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := []Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	db.Find(&bookings)
	for _, booking := range bookings {

		s, err := strconv.Atoi(key)
		if err == nil {
			if booking.Id == s {
				fmt.Println(booking)
				fmt.Println("Endpoint Hit: Booking No:", key)
				json.NewEncoder(w).Encode(booking)
			}
		}
	}
}

func handleRequests() {
	log.Println("Starting server at  :8000 ")
	log.Println("Quit the server with CONTROL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", returnAllBookings).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	db, err = gorm.Open("mysql", "root:newpassword@tcp(127.0.0.1:3306)/febin?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Booking{})
	handleRequests()
}
