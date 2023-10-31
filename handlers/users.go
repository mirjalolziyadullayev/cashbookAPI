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
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	}
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	var usersData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &usersData)

	for i := 0; i < len(usersData); i++ {
		var accountsData []models.Account
		byteData,_ := os.ReadFile("db/accounts.json")
		json.Unmarshal(byteData, &accountsData)

		for j := 0; j < len(accountsData); j++ {
			if usersData[i].ID ==  accountsData[j].UserID {
				usersData[i].Account = append(usersData[i].Account, accountsData[j])

				var transactionsData []models.Transaction
				byteData,_ := os.ReadFile("db/transactions.json")
				json.Unmarshal(byteData, &transactionsData)

				for k := 0; k < len(transactionsData); k++ {
					if accountsData[j].ID == transactionsData[k].AccountID {
						accountsData[i].Transactions = append(accountsData[i].Transactions, transactionsData[k])
					}
				}
			}
		}
	}
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	var usersData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &usersData)

	newUser.ID = len(usersData)+1
	newUser.Since = time.Now()
	usersData = append(usersData, newUser)

	result,_ := json.Marshal(usersData)
	os.WriteFile("db/users.json", result, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Println("user created!")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.User
	json.NewDecoder(r.Body).Decode(&updateUser)

	var usersData []models.User
	byteData,_ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &usersData)

	var userFound bool
	for i := 0; i < len(usersData); i++ {
		if usersData[i].ID == updateUser.ID {
			usersData[i].Firstname = updateUser.Firstname
			usersData[i].Lastname = updateUser.Lastname
			usersData[i].Edit = time.Now()

			userFound = true
			break
		}
 	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("User with such an ID not found.")
		return
	}
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser models.User
	json.NewDecoder(r.Body).Decode(&deleteUser)

	var usersData []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &usersData)

	var userFound bool

	for i := 0; i < len(usersData); i++ {
		if usersData[i].ID == deleteUser.ID {
			usersData = append(usersData[:i],usersData[i+1:]... )
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("User with such an ID not found.")
		return
	}

	result,_ := json.Marshal(usersData)
	os.WriteFile("db/users.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("user deleted!")
}