package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleOpenLogin(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{
	})
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