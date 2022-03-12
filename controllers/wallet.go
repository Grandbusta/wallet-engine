package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"wallet-engine/config"
	"wallet-engine/models"
	"wallet-engine/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type updatedWallet struct {
	models.Wallet
	Added_amount   float64 `json:"added_amount"`
	Removed_amount float64 `json:"removed_amount"`
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wal models.Wallet
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&wal)
	if err != nil {
		fmt.Println(err)
	}
	exist := db.Where("email=?", wal.Email).First(&models.Wallet{})
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to create wallet")
		return
	}
	if exist.RowsAffected > 0 {
		utils.RespondWithError(w, http.StatusConflict, "email already exist")
		return
	}
	wallet := utils.GenerateWallet(wal.Email)
	newWallet := models.Wallet{Email: wal.Email, Wallet_address: wallet}
	result := db.Create(&newWallet)
	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to create wallet")
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"data":   map[string]interface{}{"data": newWallet},
	})
}

func DebitWallet(w http.ResponseWriter, r *http.Request) {
	var wal models.Wallet
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&wal)
	if err != nil {
		fmt.Println(err)
	}
	if wal.Wallet_address <= 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "wallet address not present in body")
		return
	}
	if wal.Amount <= 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, "amount must be greater than 0")
		return
	}
	exist := db.Where("wallet_address=?", wal.Wallet_address).First(&models.Wallet{})
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	if exist.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusConflict, "wallet not found")
		return
	}
	var ch models.Wallet
	db.Where("wallet_address=?", wal.Wallet_address).First(&ch)
	if wal.Amount > ch.Amount {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "insufficient funds")
		return
	}
	updErr := db.Model(&wal).Where("wallet_address=?", wal.Wallet_address).UpdateColumn("amount", gorm.Expr("amount-?", wal.Amount)).Error
	if updErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	var updated models.Wallet
	db.Where("wallet_address=?", wal.Wallet_address).First(&updated)
	data := updatedWallet{Removed_amount: wal.Amount, Wallet: updated}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   map[string]interface{}{"data": data},
	})
}

func CreditWallet(w http.ResponseWriter, r *http.Request) {
	var wal models.Wallet
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&wal)
	if err != nil {
		fmt.Println(err)
	}
	if wal.Wallet_address <= 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "wallet address not present in body")
		return
	}
	if wal.Amount <= 0 {
		utils.RespondWithError(w, http.StatusInternalServerError, "amount must be greater than 0")
		return
	}
	exist := db.Where("wallet_address=?", wal.Wallet_address).First(&models.Wallet{})
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	if exist.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusConflict, "wallet not found")
		return
	}
	var updated models.Wallet
	updErr := db.Model(&wal).Where("wallet_address=?", wal.Wallet_address).UpdateColumn("amount", gorm.Expr("amount+?", wal.Amount)).Error
	if updErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	db.Where("wallet_address=?", wal.Wallet_address).First(&updated)
	data := updatedWallet{Added_amount: wal.Amount, Wallet: updated}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   map[string]interface{}{"data": data},
	})
}

func ActivateWallet(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	vars := mux.Vars(r)
	wallet := vars["wallet"]
	exist := db.Where("wallet_address=?", wallet).First(&models.Wallet{})
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	if exist.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusConflict, "wallet not found")
		return
	}
	var updated models.Wallet
	db.Model(&models.Wallet{}).Where("wallet_address=?", wallet).Update("is_active", true)
	db.Where("wallet_address=?", wallet).First(&updated)
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   map[string]interface{}{"data": updated},
	})
}

func DeactivateWallet(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	vars := mux.Vars(r)
	wallet := vars["wallet"]
	exist := db.Where("wallet_address=?", wallet).First(&models.Wallet{})
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	if exist.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusConflict, "wallet not found")
		return
	}
	var updated models.Wallet
	db.Model(&models.Wallet{}).Where("wallet_address=?", wallet).Update("is_active", false)
	db.Where("wallet_address=?", wallet).First(&updated)
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   map[string]interface{}{"data": updated},
	})
}
