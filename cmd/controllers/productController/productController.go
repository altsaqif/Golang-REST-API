package productController

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/altsaqif/go-restapi-mux/cmd/configs"
	"github.com/altsaqif/go-restapi-mux/cmd/helpers"
	"github.com/altsaqif/go-restapi-mux/cmd/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	Response = helpers.Response
)

func Products(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	// userID, ok := r.Context().Value("userID").(uint)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	if err := configs.DB.Find(&products).Error; err != nil {

		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Product Tidak Ditemukan",
			}
			Response(w, http.StatusNotFound, response, nil)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			Response(w, http.StatusInternalServerError, response, nil)
			return
		}
	}

	response := map[string][]models.Product{
		"message": products,
	}
	Response(w, http.StatusOK, response, nil)
}

func Product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusBadRequest, response, nil)
	}

	// userID, ok := r.Context().Value("userID").(uint)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	idUint := uint(id)

	var product models.Product
	if err := configs.DB.Where("id = ?", idUint).First(&product).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Product Tidak Ditemukan",
			}
			Response(w, http.StatusNotFound, response, nil)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			Response(w, http.StatusInternalServerError, response, nil)
			return
		}

	}

	// if product.UserID != userID {
	// 	response := map[string]string{
	// 		"message": "Unauthorized",
	// 	}
	// 	Response(w, http.StatusNotFound, response, nil)
	// 	return
	// }

	response := map[string]models.Product{
		"message": product,
	}
	Response(w, http.StatusOK, response, nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Mengambil Inputan JSON Dari Body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusInternalServerError, response, nil)
		return
	}
	defer r.Body.Close()

	// userID, ok := r.Context().Value("userID").(uint)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// product.UserID = userID

	// Membuat Data Product Baru
	if err := configs.DB.Create(&product).Error; err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusBadRequest, response, nil)
		return
	}

	response := map[string]models.Product{
		"message": product,
	}
	Response(w, http.StatusOK, response, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusBadRequest, response, nil)
	}

	// Mengambil Inputan JSON Dari Body
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusInternalServerError, response, nil)
		return
	}
	defer r.Body.Close()

	// Merubah Data Product
	if configs.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		response := map[string]string{
			"message": "Product not found",
		}
		Response(w, http.StatusNotFound, response, nil)
		return
	}

	// userID, ok := r.Context().Value("userID").(uint)
	// if !ok || product.UserID != userID {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	product.Model.ID = uint(id)
	// product.UserID = userID

	response := map[string]models.Product{
		"message": product,
	}
	Response(w, http.StatusOK, response, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Mengambil Inputan JSON Dari Body
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		Response(w, http.StatusNotFound, response, nil)
		return
	}

	defer r.Body.Close()

	var product models.Product

	// userID, ok := r.Context().Value("userID").(uint)
	// if !ok || product.UserID != userID {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	if configs.DB.Delete(&product, input["id"]).RowsAffected == 0 {
		response := map[string]string{
			"message": "Product not found",
		}
		Response(w, http.StatusBadRequest, response, nil)
		return
	}

	response := map[string]string{
		"message": "Product Berhasil Dihapus",
	}
	Response(w, http.StatusOK, response, nil)
}
