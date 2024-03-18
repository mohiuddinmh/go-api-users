package handlers

import (
	"encoding/json"
	"net/http"

	"lets-go/database"
	"lets-go/repositories"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userRepository := repositories.NewUserRepository(database.Db)
	users, err := userRepository.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
