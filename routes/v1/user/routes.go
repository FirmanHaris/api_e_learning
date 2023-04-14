// nama package sesuai nama folder
package user

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/handler"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type V1UserRouteHandler struct {
	userHandler handler.UserHandler
	echo        *echo.Group
	context     context.Context
	database    *mongo.Database
}

func NewV1UserRouteHandler(user handler.UserHandler, echo *echo.Group, ctx context.Context, database *mongo.Database) V1UserRouteHandler {
	return V1UserRouteHandler{
		userHandler: user,
		echo:        echo,
		context:     ctx,
		database:    database,
	}
}

func (b *V1UserRouteHandler) V1UserRouteHandler() {
	userApi := b.echo.Group("/user")

	userApi.GET("/all", b.userHandler.GetAllUser)
	userApi.GET("/id", b.userHandler.GetUserById)
}
