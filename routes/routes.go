// nama package sesuai nama folder
package role

import (
	"context"

	http "github.com/FirmanHaris/api_e_learning/http"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteHandler struct {
	roleHandler http.RoleHandler
	userHandler http.UserHandler
	echo        *echo.Group
	context     context.Context
	database    *mongo.Database
}

func NewRouteHandler(
	role http.RoleHandler,
	user http.UserHandler,
	echo *echo.Group,
	ctx context.Context,
	database *mongo.Database,
) RouteHandler {
	return RouteHandler{
		roleHandler: role,
		userHandler: user,
		echo:        echo,
		context:     ctx,
		database:    database,
	}
}

func (b *RouteHandler) Routes() {

	roleApi := b.echo.Group("/role")
	roleApi.GET("/all", b.roleHandler.GetAllRole)

	userApi := b.echo.Group("/user")
	userApi.GET("/all", b.userHandler.GetAllUser)
	userApi.GET("/id", b.userHandler.GetUserById)
}
