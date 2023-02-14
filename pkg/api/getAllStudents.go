package api

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	controller "student_portal/pkg/controller/student"
)

func GetAllStudents(c *fiber.Ctx) error {

	index, e := strconv.Atoi(c.Query("index", "-1"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.Query("limit", "-1"))
	if e != nil {
		limit = -1
	}

	data, err := controller.GetAllStudents(index, limit)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(data)

}
