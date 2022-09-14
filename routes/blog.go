package routes

import (
	"blog/controllers"
	"blog/middlewares"

	"github.com/labstack/echo/v4"
)

func BlogRoutes(group *echo.Group) {
	group.GET("", controllers.GetApprovedBlogs)
	group.GET("/false", controllers.GetPendingBlogs, middlewares.IsLoggedIn, middlewares.Restricted)
	group.PATCH("/:id", controllers.UpdateBlog, middlewares.IsLoggedIn)
	group.POST("", controllers.PostBlog, middlewares.IsLoggedIn)
}
