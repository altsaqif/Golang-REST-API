package homeController

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]any)
	data["Nama"] = "Al Tsaqif Nugraha Ahmad"
	data["Mata Kuliah"] = "Skripsi"

	// Parsing Map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Parsing JSON to String
	dataString := string(jsonData)

	fmt.Fprintf(w, "Welcome to my API IBU SANTI RAHAYU..\n"+"\n"+dataString)
}
