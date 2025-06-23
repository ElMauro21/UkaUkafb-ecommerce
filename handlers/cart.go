package handlers

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenCart(c *gin.Context){
	view.Render(c,http.StatusOK,"cart.html",gin.H{
		"title": "My cart",
	})
}