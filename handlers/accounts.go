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
		getAccounts(w, r)
	case "POST":
		createAccount(w, r)
	case "PUT":
		updateAccount(w, r)
	case "DELETE":
		deleteAccount(w, r)
	}
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	var accountsData []models.Account
	byteData,_ := os.ReadFile("db/accounts.json")
	json.Unmarshal(byteData, &accountsData)

	for i := 0; i < len(accountsData); i++ {
		var transactionsData []models.Transaction
		byteData,_ := os.ReadFile("db/transactions.json")
		json.Unmarshal(byteData, &transactionsData)

		for j := 0; j < len(transactionsData); j++ {
			if transactionsData[j].AccountID == accountsData[i].ID {
				accountsData[i].Transactions = append(accountsData[i].Transactions, transactionsData[j])
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accountsData)
}
func createAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount models.Account
	json.NewDecoder(r.Body).Decode(&newAccount)

	var accountsData []models.Account
	byteData, _ := os.ReadFile("db/accounts.json")
	json.Unmarshal(byteData, &accountsData)

	var userFound bool

	for i := 0; i < len(accountsData); i++ {
		var userData []models.User
		byteData, _ := os.ReadFile("db/users.json")
		json.Unmarshal(byteData, &userData)

		for j := 0; j < len(userData); j++ {
			if userData[j].ID == newAccount.UserID {
				newAccount.ID = len(userData)+1
				newAccount.Since = time.Now()
				accountsData = append(accountsData, newAccount)
				userFound = true
				break
			}
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("transaction with such an UserID not found")
		return
	}

	result, _ := json.Marshal(accountsData)
	os.WriteFile("db/accounts.json", result, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Println("account created!")
}
func updateAccount(w http.ResponseWriter, r *http.Request) {
	var updateAccount models.Account
	json.NewDecoder(r.Body).Decode(&updateAccount)

	var accountsData []models.Account
	byteData, _ := os.ReadFile("db/accounts.json")
	json.Unmarshal(byteData, &accountsData)

	var accountFound bool
	var userFound bool 

	for i := 0; i < len(accountsData); i++ {
		if accountsData[i].ID == updateAccount.ID {
			var userData []models.User
			byteData, _ := os.ReadFile("db/users.json")
			json.Unmarshal(byteData, &userData)

			for j := 0; j < len(userData); j++ {
				if userData[i].ID == updateAccount.UserID {
					accountsData[i].Name = updateAccount.Name
					accountsData[i].Edited = time.Now()

					userFound = true
					break
				}
			}
			if !userFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Account with such an UserID not found")
				return
			}
			accountFound = true
			break
		}
	}
	if !accountFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Account with such an ID not found")
		return
	}

	result, _ := json.Marshal(accountsData)
	os.WriteFile("db/accounts.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("account updated!")
}
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	var deleteAccount models.Account
	json.NewDecoder(r.Body).Decode(&deleteAccount)

	var accountsData []models.Account
	byteData, _ := os.ReadFile("db/accounts.json")
	json.Unmarshal(byteData, &accountsData)

	var accountFound bool
	var userFound bool

	for i := 0; i < len(accountsData); i++ {
		if accountsData[i].ID == deleteAccount.ID {
			var userData []models.User
			byteData, _ := os.ReadFile("db/users.json")
			json.Unmarshal(byteData, &userData)

			for j := 0; j < len(userData); j++ {
				if userData[j].ID == deleteAccount.UserID {
					accountsData = append(accountsData[:i], accountsData[i+1:]... )
					userFound = true
					break
				}
			}
			if !userFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Account with such an UserID not found")
				return
			}
			accountFound = true
			break
		}
	}
	if !accountFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Account with such an ID not found")
		return
	}

	result, _ := json.Marshal(accountsData)
	os.WriteFile("db/users.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("account deleted.")
}