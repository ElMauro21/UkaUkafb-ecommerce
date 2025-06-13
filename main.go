package main

import (
	"net/http"

	"github.com/ElMauro21/UkaUkafb/database"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
  r := gin.Default()

  db := database.InitDB("./uka.db")
  defer db.Close()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })


  r.Run() 
}
