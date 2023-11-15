package handlers

import (
	"cashbookAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("sent GET request to /account on", time.Now().Format(time.RFC822))
		getAccounts(w, r)
	case "POST":
		fmt.Println("sent POST request to /accounts on", time.Now().Format(time.RFC822))
		createAccount(w, r)
	case "PUT":
		fmt.Println("sent PUT request to /accounts on", time.Now().Format(time.RFC822))
		updateAccount(w, r)
	case "DELETE":
		fmt.Println("sent DELETE request to /accounts on", time.Now().Format(time.RFC822))
		deleteAccount(w, r)
	}
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var Account []models.Account

	for i := 0; i < len(allUser); i++ {
		for j := 0; j < len(allUser[i].Account); j++ {
			Account = append(Account, allUser[i].Account[j])
		}
	}
	
	json.NewEncoder(w).Encode(Account)
}
func createAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount models.Account
	json.NewDecoder(r.Body).Decode(&newAccount)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userFound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == newAccount.UserID {
			newAccount.ID = len(allUser[i].Account) + 1
			newAccount.Created = time.Now()
			newAccount.Edited = time.Now()

			allUser[i].Account = append(allUser[i].Account, newAccount)
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user with such an UserID not found!")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "account created!")
}
func updateAccount(w http.ResponseWriter, r *http.Request) {
	var updateAccount models.Account
	json.NewDecoder(r.Body).Decode(&updateAccount)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool
	var accountfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == updateAccount.UserID {

			for j := 0; j < len(allUser[i].Account); j++ {
				if allUser[i].Account[j].ID == updateAccount.ID {
					allUser[i].Account[j].Name = updateAccount.Name
					allUser[i].Account[j].Balance = updateAccount.Balance
					allUser[i].Account[j].Edited = time.Now()

					accountfound = true
					break
				}
			}
			if !accountfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "account with such an ID not found!")
				return
			}
			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user with such an UserID not found!")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "account updated!")
}
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	var deleteAccount models.Account
	json.NewDecoder(r.Body).Decode(&deleteAccount)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool
	var accountfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == deleteAccount.UserID {

			for j := 0; j < len(allUser[i].Account); j++ {
				if allUser[i].Account[j].ID == deleteAccount.ID {
					allUser[i].Account = append(allUser[i].Account[:j], allUser[i].Account[j+1:]...)

					accountfound = true
					break
				}
			}
			if !accountfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "account with such an ID not found!")
				return
			}
			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user with such an UserID not found!")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "account deleted!")
}
