package server

import (
	"github.com/labstack/echo/v4"
	"os"
)

func New() (*echo.Echo, error) {
	e := echo.New()
	entClient, err := newEntClient(os.Getenv("DATA_SOURCE"))
	if err != nil {
		return nil, err
	}
	controllers := newControllers(entClient)

	e.GET("/api/users", controllers.user.GetAll())
	e.POST("/api/users", controllers.user.Create())
	e.GET("/api/users/:id", controllers.user.GetById())
	e.PUT("/api/users/:id", controllers.user.Update())
	e.DELETE("/api/users/:id", controllers.user.Delete())

	return e, nil
}
