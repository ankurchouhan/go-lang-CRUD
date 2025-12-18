package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* ---------------- CREATE USER ---------------- */

func createUser(w http.ResponseWriter, r *http.Request) {
	users, _ := readUsersFromFile()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// auto-increment ID
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1

	users = append(users, user)
	writeUsersToFile(users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

/* ---------------- GET ALL USERS ---------------- */

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := readUsersFromFile()
	if err != nil {
		users = []User{}
	}

	// IMPORTANT: never return null
	if users == nil {
		users = []User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

/* ---------------- GET USER BY ID ---------------- */

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	users, _ := readUsersFromFile()

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

/* ---------------- UPDATE USER ---------------- */

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	users, _ := readUsersFromFile()

	for i := range users {
		if users[i].ID == id {
			if err := json.NewDecoder(r.Body).Decode(&users[i]); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			users[i].ID = id
			writeUsersToFile(users)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

/* ---------------- DELETE USER ---------------- */

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	users, _ := readUsersFromFile()

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			writeUsersToFile(users)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
