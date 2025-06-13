package main

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/database"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
  r := gin.Default()

  // Load the templates 
  r.LoadHTMLGlob("templates/*.html")
  // Serve any static file
  r.Static("/static","./static")

  db := database.InitDB("./uka.db")
  defer db.Close()

  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusAccepted, "index.html" , gin.H{})
  })

  r.Run() 
}
