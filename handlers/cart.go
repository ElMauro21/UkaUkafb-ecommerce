package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

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

	cartID,err := cart.GetCartID(c,db)
	if err != nil {
		c.String(http.StatusInternalServerError, "No se ha encontrado carrito de compras.")
		return
	}

	var currentQuantity int
	err = db.QueryRow(`SELECT quantity FROM cart_items WHERE cart_id = ? AND product_id = ?`,cartID,prodID).Scan(&currentQuantity)
	
	if errors.Is(err, sql.ErrNoRows) {
		_,err := db.Exec(`INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)`, cartID,prodID,qty)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido añadir producto.")
			return
		}
	}else if err != nil {
		c.String(http.StatusInternalServerError, "Error al buscar el producto.")
	}else {
		_,err := db.Exec(`UPDATE cart_items SET quantity = quantity + ? WHERE cart_id = ? AND product_id = ?`,qty,cartID,prodID)
		if err != nil {
			c.String(http.StatusInternalServerError, "No se ha podido actualizar producto.")
			return
		}
	}

}

