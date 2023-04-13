package app

import (
	"context"

	v1 "github.com/FirmanHaris/api_e_learning/app/v1"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouteHandler(ctx context.Context, echo *echo.Echo, mongo *mongo.Database) {
	api := echo.Group("/api").Group("/v1")
	v1.RouteHandler(ctx, api, mongo)
}
