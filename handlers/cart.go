package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/cart"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenCart(c *gin.Context){
	view.Render(c,http.StatusOK,"cart.html",gin.H{
		"title": "My cart",
	})
}

func HandleAddToCart(c *gin.Context, db *sql.DB){
	cart.CreateCart(c,db)
	
	
}

