package validator

import (
	"fmt"
	"net/http"

	"github.com/FirmanHaris/api_e_learning/utils/s"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		c.Logger().Error(report)
		s.AbortWithMessageStatus(c, report.Code, fmt.Sprintf("%s", report.Message))
		// c.JSON(report.Code, report)
	}
}

func InitValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}
