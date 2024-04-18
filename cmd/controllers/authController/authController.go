package authController

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/altsaqif/go-restapi-mux/cmd/configs"
	"github.com/altsaqif/go-restapi-mux/cmd/helpers"
	"github.com/altsaqif/go-restapi-mux/cmd/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Mengambil Inputan JSON Dari Body
	var Login models.Login

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&Login); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helpers.Response(w, http.StatusInternalServerError, response, nil)
		return
	}
	defer r.Body.Close()

	// Ambil Data User Berdasarkan Email
	var user models.Users
	if err := configs.DB.Where("email = ?", Login.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Wrong Email",
			}
			helpers.Response(w, http.StatusBadRequest, response, nil)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helpers.Response(w, http.StatusInternalServerError, response, nil)
			return
		}
	}

	// Cek Apakah Password Valid
	if err := helpers.VerifyPassword(user.Password, Login.Password); err != nil {
		response := map[string]string{
			"message": "Wrong Password",
		}
		helpers.Response(w, http.StatusBadRequest, response, nil)
		return
	}

	// Proses Pembuatan Token JWT
	token, err := helpers.CreateToken(&user)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helpers.Response(w, http.StatusInternalServerError, response, nil)
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
		"message": "Successfully Login",
	}
	helpers.Response(w, http.StatusOK, response, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Mengambil Inputan JSON Dari Body
	var register models.Register

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&register); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helpers.Response(w, http.StatusInternalServerError, response, nil)
		return
	}
	defer r.Body.Close()

	// Pengecekan Password Sama Dengan Password Confirm
	if register.Password != register.PasswordConfirm {
		response := map[string]string{
			"message": "Password not match",
		}
		helpers.Response(w, http.StatusBadRequest, response, nil)
		return
	}

	// Hash Password Menggunakan bcrypt
	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helpers.Response(w, http.StatusInternalServerError, response, nil)
		return
	}

	user := models.Users{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	// Cara ke 1 Validasi Email Harus Unik
	if err := configs.DB.Create(&user).Error; err != nil {
		// Pengecekan Email Tidak Boleh Duplikat
		var dbError *mysql.MySQLError

		// Menggunakan error number 1062
		// Merupakan kode spesifik MySQL untuk kesalahan unik terkait dengan duplikat kunci
		if errors.As(err, &dbError) && dbError.Number == 1062 {
			response := map[string]string{
				"message": "Email is already available in the database",
			}
			helpers.Response(w, http.StatusBadRequest, response, nil)
			return
		} else {
			response := map[string]string{
				"message": err.Error(),
			}
			helpers.Response(w, http.StatusInternalServerError, response, nil)
			return
		}
	}

	response := map[string]string{
		"message": "Register Successfully",
	}
	helpers.Response(w, http.StatusCreated, response, nil)
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
	helpers.Response(w, http.StatusOK, response, nil)
}
