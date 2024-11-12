package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-embed-go-web/models"
	"github.com/go-embed-go-web/repository"
	"github.com/go-embed-go-web/service"
)

// Login function
func Login(db *sql.DB) {
	users := models.User{}
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Pastikan file ditutup setelah selesai

	// Decode JSON dari file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Inisialisasi repository dan service
	repo := repository.NewUserRepository(db)
	userService := service.CreateService(repo) // Pastikan nama fungsi sesuai

	// Panggil LoginService pada userService
	user, err := userService.LoginService(users)
	if err != nil {
		fmt.Println("Error:", err)
		response := models.Response{
			StatusCode: 404,
			Message:    "Account not found",
			Data:       nil,
		}
		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		response := models.Response{
			StatusCode: 200,
			Message:    "Login success",
			Data:       user,
		}
		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	}
}
