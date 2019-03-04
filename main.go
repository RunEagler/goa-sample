//go:generate goagen bootstrap -d GolandProject/goa-sample/src/design

package main

import (
	"GolandProject/goa-sample/app"
	"GolandProject/goa-sample/controllers"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("User CRUD")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user" controller
	c := controllers.NewUserController(service)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
