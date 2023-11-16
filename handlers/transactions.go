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
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool
	var accountfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == newTransaction.UserID {
			for j := 0; j < len(allUser[i].Account); j++ {
				if allUser[i].Account[j].ID == newTransaction.AccountID {
					newTransaction.ID = len(allUser[i].Account[j].Transactions) + 1

					if newTransaction.TransactionType == "income" {
						allUser[i].Account[j].Balance += newTransaction.Value
					} else if newTransaction.TransactionType == "outcome" {
						allUser[i].Account[j].Balance -= newTransaction.Value
					}

					newTransaction.Edited = time.Now()
					newTransaction.Done = time.Now()

					allUser[i].Account[j].Transactions = append(allUser[i].Account[j].Transactions, newTransaction)
					accountfound = true
					break
				}
			}
			if !accountfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "account with such an AccountID not found")
				return
			}
			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User with such an UserID not found")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "transaction created!")
}
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	var updateTransaction models.Transactions
	json.NewDecoder(r.Body).Decode(&updateTransaction)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool
	var accountfound bool
	var transactionfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == updateTransaction.UserID {
			for j := 0; j < len(allUser[i].Account); j++ {
				if allUser[i].Account[j].ID == updateTransaction.AccountID {
					for k := 0; k < len(allUser[i].Account[j].Transactions); k++ {
						if allUser[i].Account[j].Transactions[k].ID == updateTransaction.ID {

							if updateTransaction.TransactionType == "income" {
								allUser[i].Account[j].Balance += updateTransaction.Value
							} else if updateTransaction.TransactionType == "outcome" {
								allUser[i].Account[j].Balance -= updateTransaction.Value
							}

							allUser[i].Account[j].Transactions[k].Name = updateTransaction.Name
							allUser[i].Account[j].Transactions[k].Value = updateTransaction.Value
							allUser[i].Account[j].Transactions[k].TransactionType = updateTransaction.TransactionType
							allUser[i].Account[j].Transactions[k].Edited = time.Now()

							transactionfound = true
							break
						}
					}
					if !transactionfound {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprint(w, "transaction with such an ID not found")
						return
					}
					accountfound = true
					break
				}
			}
			if !accountfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "account with such an AccountID not found")
				return
			}
			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User with such and UserID not found")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "transaction updated!")
}
func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	var deleteTransaction models.Transactions
	json.NewDecoder(r.Body).Decode(&deleteTransaction)

	var allUser []models.User
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &allUser)

	var userfound bool
	var accountfound bool
	var transactionfound bool

	for i := 0; i < len(allUser); i++ {
		if allUser[i].ID == deleteTransaction.UserID {
			for j := 0; j < len(allUser[i].Account); j++ {
				if allUser[i].Account[j].ID == deleteTransaction.AccountID {
					for k := 0; k < len(allUser[i].Account[j].Transactions); k++ {
						if allUser[i].Account[j].Transactions[k].ID == deleteTransaction.ID {

							if allUser[i].Account[j].Transactions[k].TransactionType == "income" {
								allUser[i].Account[j].Balance -= allUser[i].Account[j].Transactions[k].Value
							} else if allUser[i].Account[j].Transactions[k].TransactionType == "outcome" {
								allUser[i].Account[j].Balance += allUser[i].Account[j].Transactions[k].Value
							}

							allUser[i].Account[j].Transactions = append(allUser[i].Account[j].Transactions[:k], allUser[i].Account[j].Transactions[k+1:]...)
							transactionfound = true
							break
						}
					}
					if !transactionfound {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprint(w, "Transaction with such and ID not found")
						return
					}
					accountfound = true
					break
				}
			}
			if !accountfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "Account with such and AccountID not found")
				return
			}
			userfound = true
			break
		}
	}
	if !userfound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User with such and UserID not found")
		return
	}

	result, _ := json.Marshal(allUser)
	os.WriteFile("db/all.json", result, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "transaction deleted!")
}
