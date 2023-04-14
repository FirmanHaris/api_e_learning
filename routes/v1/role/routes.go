// nama package sesuai nama folder
package role

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/role/handler"
	"github.com/FirmanHaris/api_e_learning/app/v1/role/service"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type V1RoleRouteHandler struct {
	roleService service.RoleService
	echo        *echo.Group
	context     context.Context
	database    *mongo.Database
}

func NewV1RoleRouteHandler(role service.RoleService, echo *echo.Group, ctx context.Context, database *mongo.Database) V1RoleRouteHandler {
	return V1RoleRouteHandler{
		roleService: role,
		echo:        echo,
		context:     ctx,
		database:    database,
	}
}

func (b *V1RoleRouteHandler) V1RoleRouteHandler() {
	roleApi := b.echo.Group("/role")

	httpRole := handler.NewRoleHandler(b.roleService, b.context)
	roleApi.GET("/all", httpRole.GetAllRole)
}
