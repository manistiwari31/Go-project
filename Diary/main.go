package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Diary struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Body  string `json:"body"`
}

var diary []Diary

func getdiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diary)
}

func creatediary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diary)
}

func updatediary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diary)
}
func deletediary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range diary {
		if item.ID == params["id"] {
			diary = append(diary[:index], diary[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(diary)
}

func main() {
	r := mux.NewRouter()

	diary = append(diary, Diary{ID: "1", Title: "lets do", Date: "23/02", Body: "fiowehfgiohg fjweiojt jfeio;wh hfioehqwoi"})
	diary = append(diary, Diary{ID: "2", Title: "hi in the", Body: "fiowehfgiohg fjweiojt jfeio;wh hfioehqwoi"})
	diary = append(diary, Diary{ID: "3", Title: "first day", Body: "fiowehfgiohg fjweiojt jfeio;wh hfioehqwoi"})

	r.HandleFunc("/diary", getdiary).Methods("GET")
	r.HandleFunc("/diary/{id}", getdiary).Methods("GET")
	r.HandleFunc("/diary", creatediary).Methods("POST")
	r.HandleFunc("/diary/{id}", updatediary).Methods("PUT")
	r.HandleFunc("/diary/{id}", deletediary).Methods("DELETE")

	fmt.Printf("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
