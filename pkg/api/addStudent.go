package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	controller "student_portal/pkg/controller/student"
	"student_portal/pkg/entity"
)

func AddStudent(c *fiber.Ctx) error {

	student := entity.Student{}
	body := c.Body()

	err := json.Unmarshal(body, &student)
	if err != nil {
		return c.JSON(err.Error())
	}

	data, err := controller.AddStudent(&student)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(data)

}
