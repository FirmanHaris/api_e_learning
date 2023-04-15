// nama package sesuai nama folder
package http

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/service"
	"github.com/FirmanHaris/api_e_learning/utils/s"
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
	return s.Auto(c, data, err)
}
