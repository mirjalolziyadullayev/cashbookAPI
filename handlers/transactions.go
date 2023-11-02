package handlers

import (
	"net/http"
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

}
func createTransaction(w http.ResponseWriter, r *http.Request) {

}
func updateTransaction(w http.ResponseWriter, r *http.Request) {

}
func deleteTransaction(w http.ResponseWriter, r *http.Request) {

}
