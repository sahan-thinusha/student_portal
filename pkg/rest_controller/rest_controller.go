package rest_controller

import (
	"github.com/gofiber/fiber/v2"
	"student_portal/pkg/api"
)

func FiberController(app *fiber.App) {

	app.Post("/v1/api/student", api.AddStudent)
	app.Put("/v1/api/student", api.UpdateStudent)
	app.Get("/v1/api/student", api.GetAllStudents)

}
