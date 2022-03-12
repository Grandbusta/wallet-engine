package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wallet-engine/config"
	"wallet-engine/models"
)

type Wallet struct {
	Email string
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wal Wallet
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&wal)
	if err != nil {
		fmt.Println(err)
	}
	wallet := models.Wallet{}

	existErr := db.Model(&models.Wallet{}).First(&wallet, wal.Email).Error
	fmt.Println(existErr)
	fmt.Println(wal)
}

func DebitWallet(w http.ResponseWriter, r *http.Request) {

}

func CreditWallet(w http.ResponseWriter, r *http.Request) {

}

func ActivateWallet(w http.ResponseWriter, r *http.Request) {

}

func DeactivateWallet(w http.ResponseWriter, r *http.Request) {

}
