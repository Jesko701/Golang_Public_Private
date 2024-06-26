package controllers

import (
	"encoding/json"
	"net/http"
	"simple_web_server/students"
)

var data = []students.Students{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func Users(w http.ResponseWriter, r *http.Request) { // * Get All Conventional
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// * convert data GO into marshal for send packet to the network
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var result []byte
	var err error

	for _, each := range data {
		if each.ID == id {
			result, err = json.Marshal(each)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
