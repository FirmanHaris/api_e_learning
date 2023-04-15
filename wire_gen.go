// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/FirmanHaris/api_e_learning/http"
	"github.com/FirmanHaris/api_e_learning/repository"
	"github.com/FirmanHaris/api_e_learning/routes"
	"github.com/FirmanHaris/api_e_learning/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitializeRouteHandler(context2 context.Context, echo2 *echo.Group, database *mongo.Database) role.RouteHandler {
	roleRepository := repository.NewRoleRepository(database)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := http.NewRoleHandler(roleService, context2)
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService, context2)
	routeHandler := role.NewRouteHandler(roleHandler, userHandler, echo2, context2, database)
	return routeHandler
}
