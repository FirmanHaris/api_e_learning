// nama package sesuai nama folder
package user

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/handler"
	"github.com/FirmanHaris/api_e_learning/app/v1/user/repository"
	"github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(ctx context.Context, e *echo.Group, database *mongo.Database) {
	collection := database.Collection("user")
	userRepo := repository.NewUserRepository(collection)
	userService := service.NewUserService(userRepo)

	httpUser := handler.UserHandler{
		User: userService,
		Ctx:  ctx,
	}
	userApi := e.Group("/user")
	userApi.GET("/all", httpUser.GetAllUser)

}
