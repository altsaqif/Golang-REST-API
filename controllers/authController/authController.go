package authController

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/altsaqif/go-restapi-gin/config"
	"github.com/altsaqif/go-restapi-gin/database"
	"github.com/altsaqif/go-restapi-gin/helper"
	"github.com/altsaqif/go-restapi-gin/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Mengambil Inputan JSON
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Ambil Data User Berdasarkan Username
	var user models.User
	if err := database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Username atau Password Salah!",
			}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	// Cek Apakah Password Valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{
			"message": "Username atau Password Salah!",
		}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// Proses Pembuatan Token JWT
	expTime := time.Now().Add(time.Minute * 10)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-restapi-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Deklarasi Algoritma Untuk Signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signed Token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Set Token ke Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{
		"message": "Login Berhasil!",
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Mengambil Inputan JSON
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Hash Password Menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// Insert ke Database
	if err := database.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{
		"message": "success",
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Hapus Token Yang Ada di Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout Berhasil!"}
	helper.ResponseJson(w, http.StatusOK, response)
}
