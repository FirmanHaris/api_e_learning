package s

import (
	"errors"

	"github.com/FirmanHaris/api_e_learning/utils/r"
	echo "github.com/labstack/echo/v4"
)

func AbortWithStatus(c echo.Context, status int, err error) error {
	result := NewResultError(err)
	result.Code = status
	println(result.ErrorMessage)
	return c.JSON(result.Code, result)
}

func AbortWithMessageStatus(c echo.Context, status int, message string) error {
	return AbortWithStatus(c, status, errors.New(message))
}

func Auto(c echo.Context, data interface{}, err r.Ex) error {
	if err != nil {
		println(err.Error())
		return Abort(c, err)
	} else {
		return ResultWithData(c, data)
	}
}

func Abort(c echo.Context, err r.Ex) error {
	result := NewResultError(err)
	println(result.ErrorMessage)
	return c.JSON(result.Code, result)
}

func AbortWithMessage(c echo.Context, message string) error {
	result := NewResultError(errors.New(message))
	println(result.ErrorMessage)
	return c.JSON(result.Code, result)
}

func ResultWithData(c echo.Context, data interface{}) error {
	result := NewResultData(data)
	return c.JSON(result.Code, result)
}
