package main

import (
	"blog/connection"
	"blog/controllers"
	"blog/middlewares"
	"blog/routes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatalf("Error loading .env file")
	}
	connection.InitDB()
	if err := middlewares.Attach(e); err != nil {
		e.Logger.Error("error occur when attaching middlewares", err)
		os.Exit(1)
	}
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	blog := e.Group("/api/v1/blogs")
	routes.BlogRoutes(blog)
	education := e.Group("/api/v1/educations")
	routes.EducationRoutes(education)
	e.Static("/api/v1/assets", "assets")
	e.POST("/api/v1/register", controllers.Registration)
	e.POST("/api/v1/login", controllers.Login)
	e.GET("/api/v1/health", controllers.Health)
	e.POST("/api/v1/imageupload", controllers.UploadImage, middlewares.IsLoggedIn)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Ok"})
	})

	go func() {
		e.Logger.Fatal(e.Start(":8000"))
	}()

	GracefulShutdown(e)
}

// server will gracefully shutdown within 5 sec
func GracefulShutdown(e *echo.Echo) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	e.Logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Info("server shutdowns gracefully")
	}
}
