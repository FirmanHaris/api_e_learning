package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/FirmanHaris/api_e_learning/app"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
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

	app.RouteHandler(ctx, e, database)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to api e-learning techcode")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
