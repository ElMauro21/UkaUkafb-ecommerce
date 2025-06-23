package cart

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCart(c *gin.Context, db *sql.DB){

	session := sessions.Default(c)
	email := session.Get("user")
	
	if email != nil{
		var userID int 
		err := db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&userID)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido encontrar usuario.")
			return
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido verificar carrito de compras.")
			return
		}
		
		if count > 0{
			return
		}

		_,err = db.Exec(`INSERT INTO carts
		(user_id) VALUES (?)`,
		userID,
		)
		if err != nil{
			c.String(http.StatusInternalServerError, "No se ha podido crear carrito.")
			return
		}
	} else {
		sessionID := session.Get("cart_session_id")
		if sessionID == nil {
			sessionID = uuid.New().String()
			session.Set("cart_session_id",sessionID)
			session.Save()
			if err := session.Save(); err != nil {
				c.String(http.StatusInternalServerError, "No se pudo guardar la sesión.")
				return
			}
		}

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM carts WHERE session_id = ?", sessionID).Scan(&count)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido verificar carrito de compras.")
			return
		}
		
		if count > 0{
			return
		}

		_,err = db.Exec(`INSERT INTO carts
		(session_id) VALUES (?)`,
		sessionID,
		)
		if err != nil{
			c.String(http.StatusInternalServerError, "No se ha podido crear carrito.")
			return
		}
	}
}