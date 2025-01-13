package main

import (
	"gc2-p3-xiowel/config"
	"gc2-p3-xiowel/grpc"
	"gc2-p3-xiowel/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"path/filepath"
)

var jwtSecret []byte

func init() {
	// Muat .env
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../.env")
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found, loading environment variables directly")
	}

	//Inisialisasi JWT Secret
	config.InitJwtSecret()

	// Validate var env
	validateEnv()
}

func validateEnv() {
	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}
}

func main() {
	// Inisialisasi database
	config.InitDB()
	defer config.CloseDB()

	go grpc.StartGRPCServer()

	// Echo server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	grpcClient := grpc.ConnectGRPCClient()
	routes.SetupRoutes(e, grpcClient)

	// Setup rute dan jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
