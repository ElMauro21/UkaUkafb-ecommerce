package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleOpenLogin(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{
	})
}

func HandleLogin(c *gin.Context, db *sql.DB){
	email := c.PostForm("log-email")
	password := c.PostForm("log-password")

	var storedHash string
	var isAdmin int

	err := db.QueryRow("SELECT password_hash, is_admin FROM users WHERE email = ?", email).Scan(&storedHash, &isAdmin)
	if err != nil {
		c.String(http.StatusBadRequest, "Wrong user or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash),[]byte(password))
	if err != nil {
		c.String(http.StatusInternalServerError, "Wrong user or password.")
		return
	}

	session := sessions.Default(c)
	session.Set("user", email)
	session.Set("admin", isAdmin == 1)
	session.Set("loggedIn", true)
	session.Save()
	c.Redirect(http.StatusSeeOther,"/")
}

func HandleLogout(c *gin.Context){
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/auth/login")
}

func HandleRegister(c *gin.Context, db *sql.DB){
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	idNumber := c.PostForm("id-number")
	phone := c.PostForm("phone")
	mail := c.PostForm("reg-mail")
	state := c.PostForm("state")
	city := c.PostForm("city")
	neigb := c.PostForm("neighborhood")
	addr := c.PostForm("address")
	pass1 := c.PostForm("reg-password1")
	pass2 := c.PostForm("reg-password2")

	if pass1 != pass2{
		c.String(http.StatusBadRequest, "Passwords do not match.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass1), bcrypt.DefaultCost)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create user.")
		return
	}

	_, err = db.Exec(`INSERT INTO users 
	(names, surnames, id_number, phone, email, state, city, neighborhood, address, password_hash) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, 
	name, surname, idNumber, phone, mail, state, city, neigb, addr, hashedPassword)
	if err != nil{
		c.String(http.StatusInternalServerError, "Can not create user.")
		return
	}
	
	c.Redirect(http.StatusSeeOther,"/auth/login")
}

func HandleCreateAdminUser(db *sql.DB){
	name := "Admin"
	email := "admin@example.com"
	password := "admin123" 

	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if count > 0{
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	_,err := db.Exec("INSERT INTO users (names, email, password_hash, is_admin) VALUES (?,?,?,1)",name, email, hashedPassword)
	if err != nil {
		log.Printf("Failed to create admin user: %v", err)
	}
}

func HandleCreateRecoveryLink(c *gin.Context, db *sql.DB){
	email:= c.PostForm("recover-email")

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil{
		c.String(http.StatusInternalServerError, "Database error.")
		return 
	}
	
	if count == 0{
		c.String(http.StatusOK, "If this email exists, a recovery email has been sent.")
		return 
	}

	token, err := generateResetToken(32)
	if err != nil {
		c.String(http.StatusInternalServerError, "Can not create reset token.")
		return
	}

	expiresAt := time.Now().Add(15 * time.Minute)
	_, err = db.Exec(`INSERT INTO password_resets 
	(email, token, expires_at) VALUES (?, ?, ?)`,
	 email,token,expiresAt)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to store reset token.")
		return
	}

	resetLink := fmt.Sprintf("http://localhost:8080/auth/recover/reset?token=%s", token)
	
	c.Set("reset_link",resetLink)
	c.Set("reset_email",email)

	if err := sendRecoveryEmail(c); err != nil {
	log.Println("Failed to send recovery email:", err)
	c.String(http.StatusInternalServerError, "Could not send recovery email.")
	return
	}

	c.Redirect(http.StatusSeeOther,"/auth/login")
}

func generateResetToken(n int) (string, error){
	bytes := make([]byte,n)
	_,err := rand.Read(bytes)
	if err != nil {
		return "",err
	}
	return hex.EncodeToString(bytes),nil
}

func sendRecoveryEmail(c *gin.Context) error {
	
	emailRaw, emailExists := c.Get("reset_email")
	resetLink, linkExists := c.Get("reset_link")
	
	if !emailExists || !linkExists{
		return fmt.Errorf("context data missing: reset_email or reset_link not set")
	}

	email, ok := emailRaw.(string)
	if !ok {
		return fmt.Errorf("email in context is not a string")
	}
	
	password := os.Getenv("SMTP_PASSWORD")
  	if password == ""{
    	log.Fatal("SMTP_PASSWORD is not set")
  	}

	var(
		smtpHost = "smtp.gmail.com"
		smtpPort = "587"
		smtpUsername = "mauro311095@gmail.com"
	)
	
	from := smtpUsername
	to := []string{email}

	subject := "Subject: Recuperar contraseña\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	
	htmlBody := fmt.Sprintf(`
	<html>
  		<body>
		<img src="https://i.postimg.cc/jjkv7qRM/temp-Imagey0-Jawd.avif" alt="Company banner" width="100" />
    		<p>Hola,</p>
    		<p>Has solicitado recuperar tu contraseña. Haz clic en el siguiente botón para restablecerla:</p>
    		<a href="%s" style="display:inline-block;padding:10px 15px;background-color:#007BFF;color:white;text-decoration:none;border-radius:5px;">Restablecer contraseña</a>
    		<p>Si no solicitaste esto, puedes ignorar este mensaje.</p>
  		</body>
	</html>
	`, resetLink)

	message := []byte(subject + mime + htmlBody )

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}

func HandleShowResetForm(c *gin.Context){
	token := c.Query("token")

	c.HTML(http.StatusOK,"reset-password.html",gin.H{
		"token": token,
	})
}

func HandleResetPassword(c *gin.Context, db *sql.DB){
	token := c.PostForm("token")
	newPassword := c.PostForm("recover-pass")

	if token == "" || newPassword == ""{
		c.String(http.StatusBadRequest, "Token and password are required.")
		return
	}

	var email string
	var expiresAt time.Time
	err := db.QueryRow("SELECT email, expires_at FROM password_resets WHERE token = ?", token).Scan(&email, &expiresAt)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid or expired token.")
		return
	}

	if time.Now().After(expiresAt) {
		c.String(http.StatusBadRequest, "Token has expired.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to update password.")
		return
	}
	
	hashedPassword := hash 

	_, err = db.Exec("UPDATE users SET password_hash = ? WHERE email = ?", hashedPassword, email)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update password.")
		return
	}

	_, _ = db.Exec("DELETE FROM password_resets WHERE token = ?", token)

	c.Redirect(http.StatusSeeOther, "/auth/login")

}
