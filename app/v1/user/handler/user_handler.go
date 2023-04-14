// nama package sesuai nama folder
package handler

import (
	"context"
	"net/http"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
	ctx         context.Context
}

func NewUserHandler(user service.UserService, ctx context.Context) UserHandler {
	return UserHandler{userService: user, ctx: ctx}
}

// htpp berisi handler untuk api
func (b *UserHandler) GetAllUser(c echo.Context) error {
	data, err := b.userService.GetAllUser(b.ctx)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, data)
}
