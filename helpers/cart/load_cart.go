package cart

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type CartItem struct{
    ProductID int
    Image string
    Name string
    Quantity int
    Price float64
    Subtotal float64
}

func LoadCartItems(c *gin.Context, db *sql.DB) ([]CartItem,error){
    
    cartID , err := GetCartID(c,db)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
		return []CartItem{}, nil
	}
	    return nil, err
    }

    rows, err := db.Query(`
    SELECT 
    cart_items.product_id,
    products.image_url,
    products.name,
    cart_items.quantity,
    products.price
    FROM 
    cart_items
    JOIN
    products ON cart_items.product_id = products.id
    WHERE 
    cart_items.cart_id = ?
    `,cartID)

    if err != nil {
        return []CartItem{},err
	}

    defer rows.Close()

    var items []CartItem

    for rows.Next() {
        var item CartItem
        err := rows.Scan(&item.ProductID, &item.Image, &item.Name, &item.Quantity, &item.Price)
        if err != nil {
            continue
        }
        item.Subtotal = item.Price * float64(item.Quantity)
        items = append(items, item)
    }

    return items,nil
}