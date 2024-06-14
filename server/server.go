package server

import (
	"github.com/labstack/echo/v4"
)

func New() (*echo.Echo, error) {
	e := echo.New()
	entClient, err := newEntClient("root:password@tcp(clean_architecture_db:3306)/clean_architecture?parseTime=true")
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
