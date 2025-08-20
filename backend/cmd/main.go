package main

import (
	"context"
	"fmt"
	"log"
	"mersinden-stockapp/internal/config"
	"mersinden-stockapp/internal/firebase_client"
	"mersinden-stockapp/internal/handlers"
	"mersinden-stockapp/internal/repositories"
	"mersinden-stockapp/internal/services"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/api/option"
)

func main() {
	//setup
	AppConfig := config.LoadConfig()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: AppConfig.ServerParams.FrontendAddress,
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		// AllowCredentials: true, // enable if you use cookies/auth headers that need credentials
	}))

	//middleware
	//rate limit

	//fbase auth
	//create options file
	opt := option.WithCredentialsFile(AppConfig.GoogleParams.GoogleApplicationCredentials)
	//get fbase instance
	ctx := context.Background()

	//disable for prod
	os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "127.0.0.1:9099")

	fbaseApp, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "mersinden-stockapp",
	}, opt)
	if err != nil {
		panic("failed to initialize firebase instance")
	}
	//get auth from fbase instance
	auth, err := fbaseApp.Auth(ctx)
	if err != nil {
		panic("failed to pull auth from firebase instance")
	}

	//create firebase_client with auth
	firebaseClient := firebase_client.NewFirebaseClient(auth)
	app.Use(firebaseClient.AuthenticateToken)

	//dependencies
	//database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", AppConfig.DBParams.Host, AppConfig.DBParams.User, AppConfig.DBParams.Password, AppConfig.DBParams.Name, AppConfig.DBParams.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to database", err)
	}
	defer close(db)

	//non-import dependencies
	productRepo := repositories.NewProductRepository(db)
	merchantRepo := repositories.NewMerchantRepository(db)
	service := services.NewServicesStruct(productRepo, merchantRepo)
	handler := handlers.NewHandler(service)

	//routes
	app.Get("/items", handler.GetItems)
	app.Post("/items", handler.CreateItem)
	app.Put("/items/:id", handler.UpdateItem)
	app.Delete("/items/:id", handler.DeleteItem)

	app.Get("/merchant/:id", handler.GetMerchantInfo)
	app.Get("/merchant", handler.GetMerchantSelf)
	app.Put("/merchant", handler.UpdateMerchantInfo)

	app.Listen(AppConfig.ServerParams.ListenPort)

}

func close(db *gorm.DB) {
	db_, err := db.DB()
	if err != nil {
		panic(err)
	}
	db_.Close()
}
