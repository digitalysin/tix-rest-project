package controller

import (
	_ "tix-rest-project/docs"
	"tix-rest-project/internal/shared"

	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Deps shared.Deps
		Post Post
	}
)

func (impl *Holder) RegisterRoutes() {
	var (
		app = impl.Deps.Echo
	)

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	// - swagger
	app.GET("/swagger/*", swagger.WrapHandler)

	app.GET("/posts", impl.Post.Find)
	app.GET("/posts/:id", impl.Post.FindById)
	app.POST("/posts", impl.Post.Create)
	app.PATCH("/posts/:id", impl.Post.Update)
	app.DELETE("/posts/:id", impl.Post.Delete)
}
