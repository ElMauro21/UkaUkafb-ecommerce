package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleOpenProfile(c *gin.Context, db *sql.DB){
	session := sessions.Default(c)
	email := session.Get("user")
	if email == nil {
		c.Redirect(http.StatusSeeOther,"/auth/login")
		return
	}

	var user struct {
        Names       string
        Surnames    string
        IDNumber    string
        Phone       string
        State       string
        City        string
        Neighborhood string
        Address     string
    }

	    err := db.QueryRow(`
        SELECT names, surnames, id_number, phone, state, city, neighborhood, address 
        FROM users WHERE email = ?`, email).
        Scan(&user.Names, &user.Surnames, &user.IDNumber, &user.Phone,
            &user.State, &user.City, &user.Neighborhood, &user.Address)

    if err != nil {
        c.String(http.StatusInternalServerError, "Error loading profile.")
        return
    }

    view.Render(c, http.StatusOK, "profile.html", gin.H{
        "title": "Perfil",
        "User":  user,
    })
}

func HandleUpdateProfile(c *gin.Context, db *sql.DB) {
	session := sessions.Default(c)
	email := session.Get("user")
	if email == nil {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	}

	name := c.PostForm("name")
	surname := c.PostForm("surname")
	idNumber := c.PostForm("id-number")
	phone := c.PostForm("phone")
	state := c.PostForm("state")
	city := c.PostForm("city")
	neighborhood := c.PostForm("neighborhood")
	address := c.PostForm("address")

	_, err := db.Exec(`
        UPDATE users
        SET names = ?, surnames = ?, id_number = ?, phone = ?, state = ?, city = ?, neighborhood = ?, address = ?
        WHERE email = ?
    `, name, surname, idNumber, phone, state, city, neighborhood, address, email)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating profile.")
		return
	}

	c.Redirect(http.StatusSeeOther, "/user/profile")
}
