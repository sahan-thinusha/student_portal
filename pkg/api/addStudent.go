package api

import (
	echo "github.com/labstack/echo/v4"
	"student_portal/pkg/entity"
)

func AddStudent(c echo.Context) (*entity.Student, error) {

	student := entity.Student{}
	if error := c.Bind(&student); error != nil {
		return nil, error
	}

	return nil, nil
}
