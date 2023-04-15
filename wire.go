//go:build wireinject
// +build wireinject

package main

import (
	"context"
	http "github.com/FirmanHaris/api_e_learning/http"
	repository "github.com/FirmanHaris/api_e_learning/repository"
	routes "github.com/FirmanHaris/api_e_learning/routes"
	service "github.com/FirmanHaris/api_e_learning/service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRouteHandler(
	context context.Context,
	echo *echo.Group,
	database *mongo.Database,
) routes.RouteHandler {
	wire.Build(
		repository.NewRoleRepository,
		service.NewRoleService,
		http.NewRoleHandler,
		repository.NewUserRepository,
		service.NewUserService,
		http.NewUserHandler,
		routes.NewRouteHandler,
	)
	return routes.RouteHandler{}
}
