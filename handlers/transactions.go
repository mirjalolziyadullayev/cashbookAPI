package handlers

import (
	"cashbookAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTransactions(w, r)
	case "POST":
		createTransaction(w, r)
	case "PUT":
		updateTransaction(w, r)
	case "DELETE":
		deleteTransaction(w, r)
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionData []models.Transaction
	byteData, _ := os.ReadFile("db/transactions.json")
	json.Unmarshal(byteData, &transactionData)

	json.NewEncoder(w).Encode(transactionData)
	w.WriteHeader(http.StatusOK)
}
func createTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.Transaction
	json.NewDecoder(r.Body).Decode(&newTransaction)

	var transactionData []models.Transaction
	byteData, _ := os.ReadFile("db/transactions.json")
	json.Unmarshal(byteData, &transactionData)

	var userFound bool

	for i := 0; i < len(transactionData); i++ {
		var userData []models.User
		byteData, _ := os.ReadFile("db/users.json")
		json.Unmarshal(byteData, &userData)

		for j := 0; j < len(userData); j++ {
			if userData[j].ID == newTransaction.UserID {
				newTransaction.ID = len(transactionData)+1
				newTransaction.Done = time.Now()
				transactionData = append(transactionData, newTransaction)
				
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

	result,_ := json.Marshal(transactionData)
	os.WriteFile("db/transactions.json",result, 0)
	
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Transaction created!")
}
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	var updateTransaction models.Transaction
	json.NewDecoder(r.Body).Decode(&updateTransaction)

	var transactionData []models.Transaction
	byteData,_ := os.ReadFile("db/transactions.json")
	json.Unmarshal(byteData, &transactionData)

	var transactionFound bool
	var userFound bool	

	for i := 0; i < len(transactionData); i++ {
		if transactionData[i].ID == updateTransaction.ID {
			var userData []models.Transaction
			byteData, _ := os.ReadFile("db/users.json")
			json.Unmarshal(byteData, &userData)

			for j := 0; j < len(userData); j++ {
				if userData[j].ID == updateTransaction.UserID {
					transactionData[i].Name = updateTransaction.Name
					transactionData[i].Value = updateTransaction.Value
					transactionData[i].Type = updateTransaction.Type
					transactionData[i].Edited = time.Now()

					userFound = true
					break
				}
			}
			if !userFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("transaction with such an UserID not found")
				return
			}
			transactionFound = true
			break
		}
	}
	if !transactionFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("transaction with such an ID not found")
		return
	}

	result, _ := json.Marshal(transactionFound)
	os.WriteFile("db/transactions.json", result, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("transaction updated!")
}
func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	var deleteTransaction models.Transaction
	json.NewDecoder(r.Body).Decode(&deleteTransaction)

	var transactionData []models.Transaction
	byteData,_ := os.ReadFile("db/transactions.json")
	json.Unmarshal(byteData, &transactionData)

	var transactionFound bool
	var userFound bool

	for i := 0; i < len(transactionData); i++ {
		if transactionData[i].ID == deleteTransaction.ID {
			var userData []models.User
			byteData, _ := os.ReadFile("db/users.json")
			json.Unmarshal(byteData, &userData)

			for j := 0; j < len(userData); j++ {
				if userData[j].ID == deleteTransaction.UserID {
					transactionData = append(transactionData[:i], transactionData[i+1:]...)
					userFound = true
					break
				}
			}
			if !userFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("transaction with such an UserID not found")
				return
			}
			transactionFound = true
			break
		}
	}
	if !transactionFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("transaction with such an ID not found")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("transaction deleted!")
}