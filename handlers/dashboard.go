package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenDashboard(c *gin.Context){

	view.Render(c,http.StatusOK,"dashboard.html",gin.H{
		"title": "Dashboard",
	})
}

func HandleAddProduct(c *gin.Context, db *sql.DB){
	name := c.PostForm("product-name")
	description := c.PostForm("product-description")
	weight := c.PostForm("product-weight")
	size := c.PostForm("product-size")
	price := c.PostForm("product-price")
	quantity := c.PostForm("product-quantity")
	image := c.PostForm("product-image")

	_, err := db.Exec(`INSERT INTO products 
	(name, description, weight, size, price, quantity, image_url) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`, 
	name, description, weight, size, price, quantity, image)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create product.")
		return
	}

	view.RenderFlash(c,http.StatusOK,"Producto creado correctamente","success")
}