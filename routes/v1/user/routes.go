// nama package sesuai nama folder
package user

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/handler"
	"github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type V1UserRouteHandler struct {
	UserService service.UserService
	Echo        *echo.Group
	Context     context.Context
	Database    *mongo.Database
}

func NewV1UserRouteHandler(user service.UserService, echo *echo.Group, ctx context.Context, database *mongo.Database) V1UserRouteHandler {
	return V1UserRouteHandler{
		UserService: user,
		Echo:        echo,
		Context:     ctx,
		Database:    database,
	}
}

func (b *V1UserRouteHandler) V1UserRouteHandler() {
	userApi := b.Echo.Group("/user")

	httpUser := handler.NewUserHandler(b.UserService, b.Context)
	userApi.GET("/all", httpUser.GetAllUser)
}
