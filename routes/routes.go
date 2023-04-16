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
	echo        *echo.Echo
	context     context.Context
	database    *mongo.Database
}

func NewRouteHandler(
	role http.RoleHandler,
	user http.UserHandler,
	echo *echo.Echo,
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
	api := b.echo.Group("/api").Group("/v1")

	roleApi := api.Group("/role")
	roleApi.GET("/all", b.roleHandler.GetAllRole)

	userApi := api.Group("/user")
	userApi.GET("/all", b.userHandler.GetAllUser)
	userApi.GET("/single/:id", b.userHandler.GetUserById)
	userApi.POST("/register", b.userHandler.RegisterUser)
	userApi.POST("/update/password", b.userHandler.UpdatePasswordUser)
}
