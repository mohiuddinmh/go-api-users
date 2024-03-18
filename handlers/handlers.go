package handlers

import (
	"encoding/json"
	"net/http"

	"lets-go/database" // Replace with the path to database.go
	"lets-go/models"   // Replace with the path to models.go
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := fetchUsersFromDB()
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

func fetchUsersFromDB() ([]models.User, error) {
	rows, err := database.Db.Query("SELECT id, last_name, first_name, dob FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.DoB); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
