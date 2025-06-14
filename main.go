package main

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/database"
	"github.com/ElMauro21/UkaUkafb/handlers"
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
    c.HTML(http.StatusOK, "index.html" , gin.H{})
  })

  r.GET("/login", handlers.HandleOpenLogin)

  r.POST("/register",func(c *gin.Context){
    handlers.HandleRegister(c,db)
  })

  r.Run() 
}
