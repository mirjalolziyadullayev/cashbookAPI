package main

import (
	"cashbookAPI/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/users", handlers.UserHandler)
	http.HandleFunc("/transactions", handlers.TransactionHandler)
	http.HandleFunc("/accounts", handlers.AccountHandler)

	port := ":8080"
	fmt.Println("server working on port ", port)

	http.ListenAndServe(port, nil)
}