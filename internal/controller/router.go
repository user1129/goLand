package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zdos/dodo_pizza/internal/repository"
)

type Router struct {
	pizzaRepo repository.PizzaDb
}

func NewRouter(pizzaRepo repository.PizzaDb) *Router {
	return &Router{
		pizzaRepo: pizzaRepo,
	}
}

func (r *Router) Init() *echo.Echo {
	eRouter := echo.New()
	eRouter.Use(middleware.Recover())
	eRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	eRouter.GET("/pizza", func(ctx echo.Context) error {
		result, err := r.pizzaRepo.GetPizzaList(ctx.Request().Context())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, result)
	})

	return eRouter
}
