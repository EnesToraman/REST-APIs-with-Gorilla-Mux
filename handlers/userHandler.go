package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"main/models"
)

// HandleFuncUsers contains all of the HandleFunc(path, f) functions of "/users" endpoint.
func HandleFuncUsers(r *mux.Router) {
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/id/{id}", GetUserByID).Methods("GET")
	r.HandleFunc("/users/email/{email}", GetUserByEmail).Methods("GET")
	r.HandleFunc("/users/username/{username}", GetUserByUsername).Methods("GET")
	r.HandleFunc("/users/isActive/{isActive}", GetUsersByActivity).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/users", DeleteUser).Methods("DELETE")
}

// GetUsers gets all the user information from "users.json" file and prints them.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	fmt.Fprint(w, string(usersBytes))
}

// CreateUser creates a user with the given information and saves it in "users.json" file.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)
	err = json.NewEncoder(w).Encode(&user)
	CheckError(err)

	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)
	users = append(users, user)
	usersByte, err := json.MarshalIndent(users, "", "	")
	CheckError(err)

	err = os.WriteFile("data/users.json", usersByte, 0644)
	CheckError(err)
}

// UpdateUser updates the user information in "users.json" file with the given ID.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)
	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)

	for i := range users {
		if users[i].ID == user.ID {
			users[i].Email = user.Email
			users[i].Username = user.Username
			users[i].FirstName = user.FirstName
			users[i].LastName = user.LastName
			users[i].IsActive = user.IsActive
		}
	}
	usersByte, err := json.MarshalIndent(users, "", "	")
	CheckError(err)

	err = os.WriteFile("data/users.json", usersByte, 0644)
	CheckError(err)
}

// DeleteUser deletes the user information from "users.json" with the given ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)
	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)

	index := -1
	for i := range users {
		if users[i].ID == user.ID {
			index = i
		}
	}
	users = append(users[:index], users[index+1:]...)
	usersByte, err := json.MarshalIndent(users, "", "	")
	CheckError(err)

	err = os.WriteFile("data/users.json", usersByte, 0644)
	CheckError(err)
}

// GetUserByID gets the user from "users.json" file by ID field provided in URL.
// URL must contain "/users/id/" followed by an ID (int) value.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	CheckError(err)

	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)
	index := -1
	for i := range users {
		if users[i].ID == id {
			index = i
		}
	}
	userByte, err := json.Marshal(users[index])
	CheckError(err)

	fmt.Fprint(w, string(userByte))
}

// GetUserByEmail gets the user from "users.json" file by Email field provided in URL.
// URL must contain "/users/email/" followed by an Email (string) value.
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	params := mux.Vars(r)
	email := params["email"]

	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)
	index := -1
	for i := range users {
		if users[i].Email == email {
			index = i
		}
	}
	userByte, err := json.Marshal(users[index])
	CheckError(err)

	fmt.Fprint(w, string(userByte))
}

// GetUserByUsername gets the user from "users.json" file by Username field provided in URL.
// URL must contain "/users/username/" followed by an Username (string) value.
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	params := mux.Vars(r)
	username := params["username"]

	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)
	index := -1
	for i := range users {
		if users[i].Username == username {
			index = i
		}
	}
	userByte, err := json.Marshal(users[index])
	CheckError(err)

	fmt.Fprint(w, string(userByte))
}

// GetUsersByActivity gets the user from "users.json" file by isActive field provided in URL.
// URL must contain "/users/isActive/" followed by an isActive (bool) value.
func GetUsersByActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	var resUsers []models.User
	params := mux.Vars(r)
	isActive := params["isActive"]
	active, err := strconv.ParseBool(isActive)
	CheckError(err)

	usersBytes, err := os.ReadFile("data/users.json")
	CheckError(err)
	json.Unmarshal(usersBytes, &users)

	if active {
		for i := range users {
			if users[i].IsActive == active {
				resUsers = append(resUsers, users[i])
			}
		}
	} else {
		for i := range users {
			if users[i].IsActive == active {
				resUsers = append(resUsers, users[i])
			}
		}
	}
	resUsersByte, err := json.Marshal(resUsers)
	CheckError(err)

	fmt.Fprint(w, string(resUsersByte))
}
