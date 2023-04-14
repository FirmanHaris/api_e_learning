// nama package sesuai nama folder
package role

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/role/handler"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type V1RoleRouteHandler struct {
	roleHandler handler.RoleHandler
	echo        *echo.Group
	context     context.Context
	database    *mongo.Database
}

func NewV1RoleRouteHandler(role handler.RoleHandler, echo *echo.Group, ctx context.Context, database *mongo.Database) V1RoleRouteHandler {
	return V1RoleRouteHandler{
		roleHandler: role,
		echo:        echo,
		context:     ctx,
		database:    database,
	}
}

func (b *V1RoleRouteHandler) V1RoleRouteHandler() {
	roleApi := b.echo.Group("/role")

	roleApi.GET("/all", b.roleHandler.GetAllRole)
}
