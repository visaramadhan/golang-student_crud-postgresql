package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-embed-go-web/database"
	"github.com/go-embed-go-web/handler"
	"github.com/go-embed-go-web/repository"
	"github.com/go-embed-go-web/service"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal connect Database!!!!", err)
	} else {
		log.Println("Success connect to database !!!")
	}
	defer db.Close()

	// Minta input endpoint sebelum memulai server HTTP
	var endpoint string
	fmt.Print("masukkan endpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	case "logout":
		handler.Logout(db)
	}

	// Mulai server HTTP setelah input endpoint diterima
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Server is running"))
		})
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Lanjutkan ke logika berbasis teks
	siswaRepo := repository.NewSiswaRepository(db)
	siswaService := service.NewSiswaService(siswaRepo)
	siswaHandler := handler.NewSiswaHandler(siswaService)

	loginService := service.NewLoginService(db)

	var action string
	for {
		if !loginService.IsLoggedIn() {
			fmt.Print("Choose action (login, exit): ")
			fmt.Scan(&action)

			switch action {
			case "login":
				var username, password string
				fmt.Print("Enter username: ")
				fmt.Scan(&username)
				fmt.Print("Enter password: ")
				fmt.Scan(&password)

				if err := loginService.Login(username, password); err != nil {
					fmt.Println("Login failed:", err)
				} else {
					fmt.Println("Login successful!")
				}
			case "exit":
				return
			default:
				fmt.Println("Invalid action.")
			}
		} else {
			fmt.Println("Masukkan Pilihan (add, update, view, delete, logout, exit):")
			fmt.Scan(&action)
			switch action {
			case "add":
				siswaHandler.AddStudent()
			case "update":
				siswaHandler.UpdateStudent()
			case "view":
				siswaHandler.ViewStudents()
			case "delete":
				siswaHandler.DeleteStudent()
			case "logout":
				loginService.Logout()
				fmt.Println("Logged out successfully.")
			case "exit":
				return
			default:
				fmt.Println("Invalid action.")
			}
		}
	}
}
