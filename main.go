package main

import (
	"context"
	"log"
	"net/http"
	"os"

	// v1 "github.com/FirmanHaris/api_e_learning/app/v1"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb+srv://elearning:Belajaraja123@cluster0.tuqkerq.mongodb.net/?retryWrites=true&w=majority"
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	e := echo.New()

	database := client.Database(os.Getenv("DATABASE_NAME"))

	api := e.Group("api").Group("v1")

	userv1 := InitializeV1UserRouteHandler(ctx, api, database)
	userv1.V1UserRouteHandler()

	rolev1 := InitializeV1RoleRouteHandler(ctx, api, database)
	rolev1.V1RoleRouteHandler()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to api e-learning techcode")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
