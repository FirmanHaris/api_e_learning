package app

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouteHandler(ctx context.Context, echo *echo.Echo, mongo *mongo.Database) {
	v1 := echo.Group("/api").Group("/v1")
	user.UserRoutes(ctx, v1, mongo)
}
