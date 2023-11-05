package handlers

import (
	"cashbookAPI/models"
	"encoding/json"
	"net/http"
	"os"
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
	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var Transactions []models.Transactions

	for i := 0; i < len(allUser); i++ {
		for j := 0; j < len(allUser[i].Account); j++ {
			for k := 0; k < len(allUser[i].Account[j].Transactions); k++ {
				Transactions = append(Transactions, allUser[i].Account[j].Transactions[k])
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Transactions)
}
func createTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.Transactions
	json.NewDecoder(r.Body).Decode(&newTransaction)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.jon")
	json.Unmarshal(userByte, &allUser)

}
func updateTransaction(w http.ResponseWriter, r *http.Request) {

}
func deleteTransaction(w http.ResponseWriter, r *http.Request) {

}
