package handlers

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/flash"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

// type Product struct{
// 	Name string
// 	Description string
// 	Weight float64
// 	Size float64
// 	Price float64
// 	Quantity int
// 	Img1 string
// 	Img2 string
// }

func HandleOpenShop(c *gin.Context){

	msg,msgType := flash.GetMessage(c)

	view.Render(c,http.StatusOK,"shop.html",gin.H{
		"title": "Tienda",
		"Message": msg,
		"MessageType": msgType,
	})
}
