package main

import (
	"fmt"
	"log"
	"net/http"
	"wallet-engine/config"
	"wallet-engine/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	//Load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env loaded")
	}
	config.ConnectDB()
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/wallet").Subrouter()
	s.HandleFunc("/create", controllers.CreateWallet).Methods("POST")
	s.HandleFunc("/debit", controllers.DebitWallet)
	s.HandleFunc("/create", controllers.CreditWallet)
	s.HandleFunc("/activate", controllers.ActivateWallet)
	s.HandleFunc("/deactivate", controllers.DeactivateWallet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	handleRequests()
}
