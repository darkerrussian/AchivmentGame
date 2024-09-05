package Server

import (
	"AchivmentGame/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("its work !"))
}

func GetChampions(w http.ResponseWriter, r *http.Request) {
	var champions []Models.Champion
	if err := DB.Find(&champions).Error; err != nil {
		http.Error(w, "Failed to retrieve champions", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(champions)
	if err != nil {
		return
	}
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/champions", GetChampions).Methods(http.MethodGet)

}
