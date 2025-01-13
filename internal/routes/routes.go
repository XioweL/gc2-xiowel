package routes

import (
	//internal "gc2-p3-xiowel/internal/middleware"
	//"gc2-p3-xiowel/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Public Routes
	//e.POST("/users/login", handler.Login)
	//e.POST("/users/register", handler.Register)

	//auth := e.Group("")
	//auth.Use(internal.CustomJwtMiddleware)
	//auth.GET("/courses", handler.GetCourses)
	//auth.POST("/enrollments", handler.EnrollCourse)
	//auth.DELETE("/enrollments/:id", handler.UnenrollCourse)
	//
	//g := e.Group("/students")
	//g.Use(internal.CustomJwtMiddleware)
	//g.GET("/me", handler.GetPersonalInfo)

	return e
}
