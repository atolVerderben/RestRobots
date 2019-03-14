package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	asciibot "github.com/mattes/go-asciibot"
)

//Robot represents the returned asciibot
type Robot struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

func getSpecificRobot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, err := asciibot.Generate(params["id"])
	if err != nil {
		rob := &Robot{
			ID:   "Not Found",
			Body: "Not Found",
		}
		json.NewEncoder(w).Encode(rob)
		return
	}

	rob := &Robot{
		ID:   params["id"],
		Body: strings.Replace(strings.Replace(asciibot.MustGenerate(params["id"]), "\n", "<br>", -1), " ", "&nbsp;", -1),
	}
	json.NewEncoder(w).Encode(rob)
}

func getRobot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rid := asciibot.RandomID()
	rob := &Robot{
		ID:   rid,
		Body: strings.Replace(strings.Replace(asciibot.MustGenerate(rid), "\n", "<br>", -1), " ", "&nbsp;", -1),
	}

	json.NewEncoder(w).Encode(rob)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/robot", getRobot).Methods("GET")
	router.HandleFunc("/robot/{id}", getSpecificRobot).Methods("GET")
	router.Handle("/", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8000", router)

	//go http.ListenAndServe(":8080", nil)
}
