package middleware

//import (
//	"context"
//	"google.golang.org/grpc"
//	"log"
//	"strings"
//
//	"gc2-p3-xiowel/config"
//	"github.com/golang-jwt/jwt/v4"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/metadata"
//	"google.golang.org/grpc/status"
//)
//
//func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	ctx, err := AuthInterceptor(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return handler(ctx, req)
//}
//
//func AuthInterceptor(ctx context.Context) (context.Context, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		log.Println("No metadata found")
//		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
//	}
//
//	token := md["authorization"]
//	if len(token) == 0 {
//		log.Println("No Authorization token found")
//		return nil, status.Errorf(codes.Unauthenticated, "Authorization token missing")
//	}
//
//	log.Println("Received Authorization token:", token[0])
//
//	tokenString := token[0]
//	if !strings.HasPrefix(tokenString, "Bearer ") {
//		log.Println("Invalid Authorization header format")
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid Authorization header format")
//	}
//
//	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
//
//	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
//
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			log.Println("Unexpected signing method")
//			return nil, status.Errorf(codes.Unauthenticated, "Unexpected signing method")
//		}
//		return config.JwtSecret, nil
//	})
//
//	if err != nil || !parsedToken.Valid {
//		log.Println("Invalid or expired token:", err)
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid or expired token")
//	}
//
//	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
//		if userID, ok := claims["user_id"].(float64); ok {
//			ctx = context.WithValue(ctx, "user_id", int(userID))
//			log.Println("User ID extracted:", userID)
//		}
//	}
//
//	return ctx, nil
//}
