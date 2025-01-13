package handler

//import (
//	"errors"
//	"gc2-p3-xiowel/config"
//	"gc2-p3-xiowel/pb"
//	"github.com/golang-jwt/jwt/v4"
//	"github.com/labstack/echo/v4"
//	"golang.org/x/crypto/bcrypt"
//	"gorm.io/gorm"
//	"log"
//	"net/http"
//	"time"
//)
//
//func Login(c echo.Context) error {
//	var req pb.LoginRequest
//	if err := c.Bind(&req); err != nil {
//		log.Println("Bind error: ", err)
//		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
//	}
//
//	if req.Username == "" || req.Password == "" {
//		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email and password are required"})
//	}
//
//	// Mencari user berdasarkan email
//	var user pb.User
//	if err := config.DB.Where("username = ?", req.Email).First(&user).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
//		}
//		log.Println("Database error:", err)
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Database Error"})
//	}
//
//	// Membandingkan password yang di-hash
//	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
//	}
//
//	// Membuat JWT token
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"user_id": user.ID,
//		"exp":     time.Now().Add(time.Hour * 5).Unix(),
//	})
//
//	tokenString, err := token.SignedString(config.JwtSecret) // Menggunakan config.JwtSecret
//	if err != nil {
//		log.Println("Token generation error:", err)
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Generate Token"})
//	}
//
//	// Mengembalikan response dengan token
//	return c.JSON(http.StatusOK, internal.LoginResponse{Token: tokenString})
//}
