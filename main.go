package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"student_portal/pkg/env"
	"student_portal/pkg/rest_controller"
)

func init() {
	if val, exist := os.LookupEnv(env.RESTPort); exist && !strings.EqualFold(val, "") {
		env.REST_Port = val
	} else {
		env.REST_Port = "8080"
	}
}

func main() {
	e := echo.New()
	rest_controller.EchoController(e)
	e.Logger.Fatal(e.Start(":" + env.REST_Port))
}
