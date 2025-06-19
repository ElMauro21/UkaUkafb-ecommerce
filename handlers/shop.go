package handlers

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/flash"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenShop(c *gin.Context){

	msg,msgType := flash.GetMessage(c)

	view.Render(c,http.StatusOK,"shop.html",gin.H{
		"title": "Tienda",
		"Message": msg,
		"MessageType": msgType,
	})
}
