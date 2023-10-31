package main

import (
	"cashbookAPI/handlers"
	"fmt"
	"net/http"
)

func main() {

	// http.HandleFunc("/", handlers.GetHome())
	http.HandleFunc("/user", handlers.UserHandler)
	http.HandleFunc("/transaction", handlers.TransactionHandler)
	http.HandleFunc("/accounts", handlers.AccountHandler)

	port := ":8080"
	fmt.Println("server working on port ", port)

	http.ListenAndServe(port, nil)
}