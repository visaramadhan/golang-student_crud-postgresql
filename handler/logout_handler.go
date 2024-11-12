package handler

import (
	"database/sql"
	"fmt"
)

// Book struct represents a book

func Logout(db *sql.DB) {
	fmt.Println("Logged out successfully.")
}
