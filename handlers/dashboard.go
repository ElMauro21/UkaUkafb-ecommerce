package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/flash"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

type Product struct {
    ID       int
    Name     string
	Description string
    Weight   int
    Size     int
    Price    float64
    Quantity int
	Image string
}

func HandleOpenDashboard(c *gin.Context,db *sql.DB){

    rows, _ := db.Query(`SELECT id, name, description, weight, size, price, quantity, image_url FROM products`)
    var products []Product

    for rows.Next() {
        var p Product
        rows.Scan(&p.ID, &p.Name, &p.Description, &p.Weight, &p.Size, &p.Price, &p.Quantity, &p.Image)
        products = append(products, p)
    }

	msg,msgType := flash.GetMessage(c)
	view.Render(c,http.StatusOK,"dashboard.html",gin.H{
		"title": "Dashboard",
		"Message": msg,
		"MessageType": msgType,
		"products": products,
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

	if name == "" || description == "" || weight == "" || size == "" || price == "" || quantity == "" || image == "" {
		view.RenderFlash(c,http.StatusOK,"Todos los campos son obligatorios","info")
		return
	}

	_, err := db.Exec(`INSERT INTO products 
	(name, description, weight, size, price, quantity, image_url) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`, 
	name, description, weight, size, price, quantity, image)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create product.")
		return
	}

	flash.SetMessage(c,"Producto creado correctamente","success")
	c.Header("HX-Redirect","/admin/dashboard")
	c.Status(http.StatusSeeOther)
}

func HandleDeleteProduct(c *gin.Context,db *sql.DB){

	productId := c.PostForm("product-id")

	if productId == "" {
		view.RenderFlash(c,http.StatusOK,"No hay producto para eliminar","info")
		return
	}

	_, err := db.Exec(`DELETE FROM products WHERE id = ?`,productId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Can not delete product.")
	}

	flash.SetMessage(c,"Producto eliminado correctamente","success")
	c.Header("HX-Redirect","/admin/dashboard")
	c.Status(http.StatusSeeOther)
}