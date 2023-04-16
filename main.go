package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/FirmanHaris/api_e_learning/utils/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb+srv://chronicle89:Belajaraja123@testing.7uzn8tb.mongodb.net/?retryWrites=true&w=majority"
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

	database := client.Database(os.Getenv("DATABASE_NAME"))

	e := echo.New()
	e.Validator = validator.InitValidator()
	e.HTTPErrorHandler = validator.InitErrorHandler()

	route := InitializeRouteHandler(ctx, e, database)
	route.Routes()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to api e-learning techcode")
	})

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
