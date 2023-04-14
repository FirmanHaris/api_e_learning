//go:build wireinject
// +build wireinject

package main

import (
	"context"
	v1 "github.com/FirmanHaris/api_e_learning/app/v1"
	userRepository "github.com/FirmanHaris/api_e_learning/app/v1/user/repository"
	userService "github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeV1RouteHandler(
	context context.Context,
	echo *echo.Echo,
	database *mongo.Database,
) v1.V1RouteHandler {
	wire.Build(
		userRepository.NewUserRepository,
		userService.NewUserService,
		v1.NewV1RouteHandler,
	)
	return v1.V1RouteHandler{}
}
