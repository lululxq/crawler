package router

import (
	"crawler/controllers"
	"github.com/kataras/iris/v12"
)

func NewRouter()  {
	app := iris.New()

	// Auth
	auth := controllers.AuthController{}
	app.Post("/login", auth.Login)

}