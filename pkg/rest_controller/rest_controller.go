package rest_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
	"student_portal/pkg/api"
)

func EchoController(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	normalRoutes := e.Group("/student_portal")

	NormalRoutes(normalRoutes)

	SwaggerAPIDoc(normalRoutes)

}

func NormalRoutes(g *echo.Group) {
	g.POST("/v1/api/student", api.AddStudent)
	g.PUT("/v1/api/student", api.UpdateStudent)
	g.GET("/v1/api/student", api.GetAllStudents)

}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
