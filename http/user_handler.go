// nama package sesuai nama folder
package http

import (
	"context"
	"net/http"

	"github.com/FirmanHaris/api_e_learning/payload"
	"github.com/FirmanHaris/api_e_learning/service"
	"github.com/FirmanHaris/api_e_learning/utils/s"
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
	return s.Auto(c, data, err)
}
func (b *UserHandler) GetUserById(c echo.Context) error {
	u := new(payload.UserGetByID)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	result, error := b.userService.GetUserById(b.ctx, u.ID)
	return s.Auto(c, result, error)
}
func (b *UserHandler) RegisterUser(c echo.Context) error {
	u := new(payload.RegisterUserPayload)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	result, err := b.userService.AddUser(b.ctx, u)
	return s.Auto(c, result, err)
}
func (b *UserHandler) UpdatePasswordUser(c echo.Context) error {
	u := new(payload.UpdatePasswordPayload)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	if err := b.userService.UpdatePasswordUser(b.ctx, u); err != nil {
		return s.AbortWithMessage(c, err.Error())
	}
	return s.AbortWithMessageStatus(c, http.StatusOK, "password updated")
}
