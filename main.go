package main

import (
	"log"
	"os"

	"github.com/ElMauro21/UkaUkafb/database"
	"github.com/ElMauro21/UkaUkafb/handlers"
	"github.com/ElMauro21/UkaUkafb/handlers/admin"
	"github.com/ElMauro21/UkaUkafb/jobs"
	"github.com/ElMauro21/UkaUkafb/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    err := godotenv.Load()
  if err !=nil {
    log.Println("No .env file found")
  }

  secret := os.Getenv("SESSION_SECRET")
  if secret == ""{
    log.Fatal("SESSION_SECRET is not set")
  }

  store := cookie.NewStore([]byte(secret))
  store.Options(sessions.Options{
	Path:     "/",        // Cookie is valid for all paths
	MaxAge:   3600,       // 1 hour
	HttpOnly: true,       // Not accessible from JS (security)
	Secure:   false,      // Use true to security
  })

  r := gin.Default()

  r.Use(sessions.Sessions("mysession",store))
  r.Use(middleware.InjectTemplateData())

  // Load the templates 
  r.LoadHTMLGlob("templates/*.html")
  // Serve any static file
  r.Static("/static","./static")

  db := database.InitDB("./uka.db")
  defer db.Close()

  jobs.JobTokenCleanup(db)

  admin.HandleCreateAdminUser(db)

  // Home routes
  r.GET("/",handlers.HandleOpenHome)

  // Auth routes
  r.GET("/auth/login", handlers.HandleOpenLogin)
  r.GET("/auth/logout",handlers.HandleLogout)
  r.POST("/auth/login",func(c *gin.Context){
    handlers.HandleLogin(c,db)
  })
  r.POST("/auth/register",func(c *gin.Context){
    handlers.HandleRegister(c,db)
  })
  r.POST("/auth/recover/initiate",func(c *gin.Context){
    handlers.HandleCreateRecoveryLink(c,db)
  })
  r.GET("/auth/recover/reset", handlers.HandleShowResetForm)
  r.POST("/auth/recover/reset", func(c *gin.Context) {
	  handlers.HandleResetPassword(c, db)
  })

  r.Run() 
}
