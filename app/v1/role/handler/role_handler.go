// nama package sesuai nama folder
package handler

import (
	"context"
	"net/http"

	"github.com/FirmanHaris/api_e_learning/app/v1/role/service"
	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService service.RoleService
	ctx         context.Context
}

func NewRoleHandler(role service.RoleService, ctx context.Context) RoleHandler {
	return RoleHandler{roleService: role, ctx: ctx}
}

// htpp berisi handler untuk api
func (b *RoleHandler) GetAllRole(c echo.Context) error {
	data, err := b.roleService.GetAllRole(b.ctx)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, data)
}
