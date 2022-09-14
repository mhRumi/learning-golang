package routes

import (
	"blog/controllers"
	"blog/middlewares"

	"github.com/labstack/echo/v4"
)

func EducationRoutes(group *echo.Group) {
	group.POST("", controllers.CreateEducation, middlewares.IsLoggedIn)
	group.GET("", controllers.GetEducationInfo, middlewares.IsLoggedIn)
}
