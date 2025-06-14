package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleOpenLogin(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{
	})
}

func HandleLogin(c *gin.Context, db *sql.DB){
	email := c.PostForm("log-email")
	password := c.PostForm("log-password")

	var storedHash string
	var isAdmin int

	err := db.QueryRow("SELECT password_hash, is_admin FROM users WHERE email = ?", email).Scan(&storedHash, &isAdmin)
	if err != nil {
		c.String(http.StatusBadRequest, "Wrong user or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash),[]byte(password))
	if err != nil {
		c.String(http.StatusInternalServerError, "Wrong user or password.")
		return
	}

	session := sessions.Default(c)
	session.Set("user", email)
	session.Set("admin", isAdmin == 1)
	session.Set("loggedIn", true)
	session.Save()
	c.Redirect(http.StatusSeeOther,"/")
}

func HandleLogout(c *gin.Context){
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

func HandleRegister(c *gin.Context, db *sql.DB){
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	idNumber := c.PostForm("id-number")
	phone := c.PostForm("phone")
	mail := c.PostForm("reg-mail")
	state := c.PostForm("state")
	city := c.PostForm("city")
	neigb := c.PostForm("neighborhood")
	addr := c.PostForm("address")
	pass1 := c.PostForm("reg-password1")
	pass2 := c.PostForm("reg-password2")

	if pass1 != pass2{
		c.String(http.StatusBadRequest, "Passwords do not match.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass1), bcrypt.DefaultCost)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create user.")
		return
	}

	_, err = db.Exec(`INSERT INTO users 
	(names, surnames, id_number, phone, email, state, city, neighborhood, address, password_hash) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, 
	name, surname, idNumber, phone, mail, state, city, neigb, addr, hashedPassword)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create user.")
		return
	}
	
	c.Redirect(http.StatusSeeOther,"/login")
}

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