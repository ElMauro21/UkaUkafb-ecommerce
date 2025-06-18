package handlers

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenDashboard(c *gin.Context){

	view.Render(c,http.StatusOK,"dashboard.html",gin.H{
		"title": "Dashboard",
	})
}