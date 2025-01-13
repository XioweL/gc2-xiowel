package middleware

//import (
//	"gc2-p3-xiowel/config"
//	"github.com/golang-jwt/jwt/v4"
//	"github.com/labstack/echo/v4"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/metadata"
//	"google.golang.org/grpc/status"
//	"log"
//	"net/http"
//	"strings"
//)
//
//// Interceptor untuk gRPC
//func AuthInterceptor(ctx context.Context) (context.Context, error) {
//	// Ambil metadata dari context
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		log.Println("No metadata found in context")
//		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized: Missing metadata")
//	}
//
//	// Ambil token dari metadata "authorization"
//	tokenArr := md["authorization"]
//	if len(tokenArr) == 0 {
//		log.Println("Missing token in metadata")
//		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized: Missing token")
//	}
//
//	// Pastikan ada "Bearer " di token
//	authHeader := tokenArr[0]
//	if !strings.HasPrefix(authHeader, "Bearer ") {
//		log.Println("Invalid token format")
//		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized: Invalid token format")
//	}
//
//	// Ekstrak token setelah "Bearer "
//	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//
//	// Parse token
//	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Pastikan metode signing adalah HMAC
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			log.Println("Unexpected signing method")
//			return nil, status.Errorf(codes.Unauthenticated, "Invalid signing method")
//		}
//		// Kembalikan secret key untuk verifikasi
//		return config.JwtSecret, nil
//	})
//
//	if err != nil || !parsedToken.Valid {
//		log.Printf("Token error: %v", err)
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid or expired token")
//	}
//
//	log.Println("Token validated successfully")
//
//	// Ekstrak claims dan simpan user_id ke context
//	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
//		userID, ok := claims["user_id"].(float64)
//		if !ok {
//			log.Println("user_id missing or invalid in claims")
//			return nil, status.Errorf(codes.Unauthenticated, "Missing or invalid user_id in token")
//		}
//
//		// Tambahkan user_id ke context
//		log.Printf("Extracted user_id: %v", int(userID))
//		ctx = context.WithValue(ctx, "user_id", int(userID))
//	} else {
//		log.Println("Invalid claims structure")
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid token claims")
//	}
//
//	return ctx, nil
//}
//
//// Middleware untuk Echo
//func AuthGrpc(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// Cek Header Authorization
//		authHeader := c.Request().Header.Get("Authorization")
//		if authHeader == "" {
//			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Authorization Header"})
//		}
//
//		// Ambil context dari request
//		ctx := c.Request().Context()
//
//		// Set metadata manual (required for gRPC context)
//		md := metadata.New(map[string]string{
//			"authorization": authHeader,
//		})
//		ctx = metadata.NewIncomingContext(ctx, md)
//
//		// Panggil AuthInterceptor
//		newCtx, err := AuthInterceptor(ctx)
//		if err != nil {
//			st, _ := status.FromError(err)
//			return c.JSON(http.StatusUnauthorized, map[string]string{"message": st.Message()})
//		}
//
//		// Pastikan untuk set context yang baru ke dalam request Echo
//		log.Println("Context before modification:", ctx)
//		c.SetRequest(c.Request().WithContext(newCtx))
//		log.Println("Context after modification:", c.Request().Context())
//
//		return next(c)
//	}
//}
