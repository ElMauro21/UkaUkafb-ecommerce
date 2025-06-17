package handlers

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenProfile(c *gin.Context){
	view.Render(c,http.StatusOK,"profile.html",gin.H{
		"title": "Perfil",
	})
}