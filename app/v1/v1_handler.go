// nama package sesuai nama folder
package v1

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/handler"
	"github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type V1RouteHandler struct {
	UserService service.UserService
	Echo        *echo.Echo
	Context     context.Context
	Database    *mongo.Database
}

func NewV1RouteHandler(user service.UserService, echo *echo.Echo, ctx context.Context, database *mongo.Database) V1RouteHandler {
	return V1RouteHandler{
		UserService: user,
		Echo:        echo,
		Context:     ctx,
		Database:    database,
	}
}

func (b *V1RouteHandler) V1RouteHandler() {
	v1 := b.Echo.Group("v1")
	userApi := v1.Group("user")

	httpUser := handler.NewUserHandler(b.UserService, b.Context)
	userApi.GET("/all", httpUser.GetAllUser)
}
