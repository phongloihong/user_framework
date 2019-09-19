package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/phongloihong/user_framework/handlers"
)

// Public collection of public route
func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")

	publicRoute.GET("/health", handlers.CheckHeath)
	publicRoute.GET("/student", handlers.GetStudents)
	publicRoute.GET("/student/id/:id", handlers.GetStudent)
	publicRoute.PATCH("/student", handlers.SearchStudent)
	publicRoute.GET("/group-student/:name", handlers.GroupStudent)

	publicRoute.POST("/user", handlers.CreateUser)
	publicRoute.POST("/auth", handlers.Auth)
}

// Staff route
func Staff(e *echo.Echo) {
	staffRoute := e.Group("/v1/staff")

	staffRoute.Use(middleware.JWT([]byte("secret")))
	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.PATCH("/student", handlers.UpdateStudent)
	staffRoute.DELETE("/student/id/:id", handlers.DeleteStudent)
}
