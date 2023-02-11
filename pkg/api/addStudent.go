package api

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
	controller "student_portal/pkg/controller/student"
	"student_portal/pkg/entity"
)

func AddStudent(c echo.Context) error {

	student := entity.Student{}
	if error := c.Bind(&student); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.AddStudent(&student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
