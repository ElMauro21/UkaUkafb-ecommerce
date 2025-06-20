package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ElMauro21/UkaUkafb/helpers/flash"
	"github.com/ElMauro21/UkaUkafb/helpers/products"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleOpenDashboard(c *gin.Context,db *sql.DB){

	session := sessions.Default(c)
	isAdmin,ok := session.Get("isAdmin").(bool)

	if !ok || !isAdmin {
		flash.SetMessage(c,"Necesitas permisos de administrador","error")
		c.Redirect(http.StatusSeeOther,"/auth/login")
    	return
	}

	products := products.LoadProducts(db)

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
	image2 := c.PostForm("product-image-two")

	if name == "" || description == "" || weight == "" || size == "" || price == "" || quantity == "" || image == "" || image2 == ""{
		view.RenderFlash(c,http.StatusOK,"Todos los campos son obligatorios","info")
		return
	}

	_, err := db.Exec(`INSERT INTO products 
	(name, description, weight, size, price, quantity, image_url, image_url_2) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, 
	name, description, weight, size, price, quantity, image, image2)
	if err != nil{
		c.String(http.StatusInternalServerError, "No se pudo crear el producto.")
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

	name := c.PostForm("product-name")
	description := c.PostForm("product-description")
	weight := c.PostForm("product-weight")
	size := c.PostForm("product-size")
	price := c.PostForm("product-price")
	quantity := c.PostForm("product-quantity")
	image := c.PostForm("product-image")
	image2 := c.PostForm("product-image-two")
	
	if name == "" || description == "" || weight == "" || size == "" || price == "" || quantity == "" || image == "" || image2 == ""{
		view.RenderFlash(c,http.StatusOK,"Todos los campos son obligatorios","info")
		return
	}

	_, err := db.Exec(`DELETE FROM products WHERE id = ?`,productId)
	if err != nil {
		c.String(http.StatusInternalServerError, "No se pudo eliminar el producto.")
	}

	flash.SetMessage(c,"Producto eliminado correctamente","success")
	c.Header("HX-Redirect","/admin/dashboard")
	c.Status(http.StatusSeeOther)
}

func HandleUpdateProduct(c *gin.Context, db *sql.DB){

	productId := c.PostForm("product-id")
	if productId == "" {
		view.RenderFlash(c,http.StatusOK,"No hay producto para actualizar","info")
		return
	}

	name := c.PostForm("product-name")
	description := c.PostForm("product-description")
	weight := c.PostForm("product-weight")
	size := c.PostForm("product-size")
	price := c.PostForm("product-price")
	quantity := c.PostForm("product-quantity")
	image := c.PostForm("product-image")
	image2 := c.PostForm("product-image-two")

	if name == "" || description == "" || weight == "" || size == "" || price == "" || quantity == "" || image == "" || image2 == ""{
		view.RenderFlash(c,http.StatusOK,"Todos los campos son obligatorios","info")
		return
	}

	_, err := db.Exec(`UPDATE products
	SET name = ?, description = ?, weight = ?, size = ?, price = ?, quantity = ?, image_url = ?, image_url_2 = ? 
	WHERE id = ?
	`, name, description, weight, size, price, quantity, image, image2, productId)

	if err != nil {
		c.String(http.StatusInternalServerError, "No se pudo actualizar el producto.")
		return
	}

	flash.SetMessage(c,"Producto actualizado correctamente","success")
	c.Header("HX-Redirect","/admin/dashboard")
	c.Status(http.StatusSeeOther)
}