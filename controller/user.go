package controller

import (
	"cleanArchitecture/useCase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	interactor *useCase.UserInteractor
}

func NewUserController(interactor *useCase.UserInteractor) *UserController {
	return &UserController{interactor: interactor}
}

func (controller *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		output, err := controller.interactor.HandleGetAll(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, output)
	}
}

func (controller *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id, _ := strconv.Atoi(c.Param("id"))

		output, err := controller.interactor.HandleGetById(ctx, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, output)

	}
}

func (controller *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var input useCase.Input
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		output, err := controller.interactor.HandleCreate(ctx, input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, output)
	}
}

func (controller *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id, _ := strconv.Atoi(c.Param("id"))

		var input useCase.Input
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		output, err := controller.interactor.HandleUpdate(ctx, id, input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, output)
	}
}

func (controller *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id, _ := strconv.Atoi(c.Param("id"))

		err := controller.interactor.HandleDelete(ctx, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
