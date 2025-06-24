package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/ElMauro21/UkaUkafb/helpers/cart"
	"github.com/ElMauro21/UkaUkafb/helpers/flash"
	"github.com/ElMauro21/UkaUkafb/helpers/view"
	"github.com/gin-gonic/gin"
)

func HandleOpenCart(c *gin.Context){
	
	msg,msgType := flash.GetMessage(c)

	view.Render(c,http.StatusOK,"cart.html",gin.H{
		"title": "My cart",
		"Message": msg,
		"MessageType": msgType,
	})
}

func HandleAddToCart(c *gin.Context, db *sql.DB){
	
	if err := cart.CreateCart(c,db); err != nil {
		c.String(http.StatusInternalServerError, "Error al crear carrito: " + err.Error())
		return
	}
	
	productID := c.PostForm("product-id")
	quantity := c.PostForm("quantity")

	prodID,err := strconv.Atoi(productID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Id de producto invalido.")
		return
	}

	qty,err := strconv.Atoi(quantity)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cantidad de producto invalido.")
		return
	}

	var stock int
	err = db.QueryRow(`SELECT quantity FROM products WHERE id = ?`,prodID).Scan(&stock)
	if err != nil {
		c.String(http.StatusInternalServerError, "No se ha podido verificar el stock del producto.")
		return
	}

	cartID,err := cart.GetCartID(c,db)
	if err != nil {
		c.String(http.StatusInternalServerError, "No se ha encontrado carrito de compras.")
		return
	}

	var currentQuantity int
	err = db.QueryRow(`SELECT quantity FROM cart_items WHERE cart_id = ? AND product_id = ?`,cartID,prodID).Scan(&currentQuantity)
	
	if errors.Is(err, sql.ErrNoRows) {
		currentQuantity = 0
		
		if currentQuantity+qty > stock {
			flash.SetMessage(c,"No se pueden añadir más productos de los que hay en stock!","error")
			c.Redirect(http.StatusSeeOther,"/shop")
			return
		}
		
		_,err := db.Exec(`INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)`, cartID,prodID,qty)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido añadir producto.")
			return
		}
	}else if err != nil {
		c.String(http.StatusInternalServerError, "Error al buscar el producto.")
	}else {

		if currentQuantity+qty > stock {
			flash.SetMessage(c,"No se pueden añadir más productos de los que hay en stock!","error")
			c.Redirect(http.StatusSeeOther,"/shop")
			return
		}

		_,err := db.Exec(`UPDATE cart_items SET quantity = quantity + ? WHERE cart_id = ? AND product_id = ?`,qty,cartID,prodID)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido actualizar producto.")
			return
		}
	}

	flash.SetMessage(c,"Producto añadido al carrito","success")
	c.Redirect(http.StatusSeeOther,"/shop")
}

