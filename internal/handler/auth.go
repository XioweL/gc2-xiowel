package handler

import (
	"errors"
	"gc2-p3-xiowel/config"
	"gc2-p3-xiowel/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	// Validate Input
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	var existingUser models.User
	if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Username already exists"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Database error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Database Error"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate password"})
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Register failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Register success"})
}

func Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Periksa Username dan Password
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Username and password are required"})
	}

	// Cek User di Database
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
		}
		log.Println("Database error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Database Error"})
	}

	// Bandingkan Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
	}

	// Buat Token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		log.Println("Token generation error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Generate Token"})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{
		Message: "Login success",
		Token:   tokenString,
	})
}
