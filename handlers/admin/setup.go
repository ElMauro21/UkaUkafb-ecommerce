package admin

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HandleCreateAdminUser(db *sql.DB){
	name := "Admin"
	email := "admin@example.com"
	password := "admin123" 

	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if count > 0{
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	_,err := db.Exec("INSERT INTO users (names, email, password_hash, is_admin) VALUES (?,?,?,1)",name, email, hashedPassword)
	if err != nil {
		log.Printf("Failed to create admin user: %v", err)
	}
}