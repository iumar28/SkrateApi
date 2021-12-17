package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage")
	fmt.Println("Endpoint Hit: Homepage")
}

type User struct {
	Id        string `json: "Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}
type Meeting struct {
	uid1 string `json: "uid1"	`
	uid2 string `json: "uid2"`
	uid3 string `json: "uid3"`
}

var Users []User
var Meetings []Meeting

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allUsers", returnAllUsers)                     //GET return all the Users
	myRouter.HandleFunc("/allMeetings", returnAllMeetings)               //GET return all the Meetings
	myRouter.HandleFunc("/newUser", createNewUser).Methods("POST")       //Add new user
	myRouter.HandleFunc("/newMeeting", createNewMeeting).Methods("POST") //Add new Meeting
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var theuser User
	json.Unmarshal(reqBody, &theuser)
	Users = append(Users, theuser)

	json.NewEncoder(w).Encode(theuser)
}
func createNewMeeting(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var themeeting Meeting
	json.Unmarshal(reqBody, &themeeting)
	Meetings = append(Meetings, themeeting)
	json.NewEncoder(w).Encode(w)
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show all users")
	json.NewEncoder(w).Encode(Users)
}
func returnAllMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show all meetings")
	json.NewEncoder(w).Encode((Meetings))
}
func main() {
	Users = []User{
		User{Id: "11", FirstName: "Umar", LastName: "Firoz", Email: "umar9897222@gmail.com"},
		User{Id: "11", FirstName: "Harish", LastName: "Khan", Email: "harishkhan@gmail.com"},
	}
	Meetings = []Meeting{
		Meeting{uid1: "1991", uid2: "1992", uid3: "1993"},
		Meeting{uid1: "1881", uid2: "1882", uid3: "1883"},
	}
	handleRequests()
}
