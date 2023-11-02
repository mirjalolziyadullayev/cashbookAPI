package handlers

import (
	"cashbookAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("sent GET request to /users on", time.Now().Format(time.RFC822))
		getUsers(w, r)
	case "POST":
		fmt.Println("sent POST request to /users on", time.Now().Format(time.RFC822))
		createUser(w, r)
	case "PUT":
		fmt.Println("sent PUT request to /users on", time.Now().Format(time.RFC822))
		updateUser(w, r)
	case "DELETE":
		fmt.Println("sent DELETE request to /users on", time.Now().Format(time.RFC822))
		deleteUser(w, r)
	}
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allUser)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	newUser.ID = len(allUser) + 1
	newUser.Created = time.Now()
	newUser.Edit = time.Now()

	allUser = append(allUser, newUser)

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created!")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.User
	json.NewDecoder(r.Body).Decode(&updateUser)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == updateUser.ID {
			allUser[i].Firstname = updateUser.Firstname
			allUser[i].Lastname = updateUser.Lastname
			allUser[i].Edit = time.Now()

			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user not found.")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "user updated")
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser models.User
	json.NewDecoder(r.Body).Decode(&deleteUser)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == deleteUser.ID {
			allUser = append(allUser[:i], allUser[i+1:]...)

			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user not found.")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "user deleted!")
}
