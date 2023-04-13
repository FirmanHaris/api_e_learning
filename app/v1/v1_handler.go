package v1

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouteHandler(ctx context.Context, echo *echo.Group, mongo *mongo.Database) {
	v1 := echo.Group("/v1")
	userRoute := v1.Group("user")
	user.UserRoutes(ctx, userRoute, mongo)

}
