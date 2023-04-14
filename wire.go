//go:build wireinject
// +build wireinject

package main

import (
	"context"
	roleRepository "github.com/FirmanHaris/api_e_learning/app/v1/role/repository"
	roleService "github.com/FirmanHaris/api_e_learning/app/v1/role/service"
	userRepository "github.com/FirmanHaris/api_e_learning/app/v1/user/repository"
	userService "github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	roleV1 "github.com/FirmanHaris/api_e_learning/routes/v1/role"
	userV1 "github.com/FirmanHaris/api_e_learning/routes/v1/user"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeV1UserRouteHandler(
	context context.Context,
	echo *echo.Group,
	database *mongo.Database,
) userV1.V1UserRouteHandler {
	wire.Build(
		userRepository.NewUserRepository,
		userService.NewUserService,
		userV1.NewV1UserRouteHandler,
	)
	return userV1.V1UserRouteHandler{}
}
func InitializeV1RoleRouteHandler(
	context context.Context,
	echo *echo.Group,
	database *mongo.Database,
) roleV1.V1RoleRouteHandler {
	wire.Build(
		roleRepository.NewRoleRepository,
		roleService.NewRoleService,
		roleV1.NewV1RoleRouteHandler,
	)
	return roleV1.V1RoleRouteHandler{}
}
