package handler

//import (
//	"github.com/golang-jwt/jwt/v4"
//	"github.com/labstack/echo/v4"
//	"net/http"
//)
//
//func BorrowBook(c echo.Context) error {
//	token, ok := c.Get("user").(*jwt.Token)
//	if !ok || token == nil {
//		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or missing token"})
//
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok || !token.Valid {
//		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
//	}
//
//	userID, ok := claims["user_id"].(string)
//	if !ok {
//		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user ID"})
//	}
//
//}
