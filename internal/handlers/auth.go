package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/thyagobr/wheretogo/internal/models"
	"github.com/thyagobr/wheretogo/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.DB.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	var loginResponse struct {
		Token string `json:"token"`
	}

	loginResponse.Token = user.Token

	respondJson(
		w,
		http.StatusOK,
		ApiResponse[struct {
			Token string `json:"token"`
		}]{
			Data: loginResponse,
		})
}
