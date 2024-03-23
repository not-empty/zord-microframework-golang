package handlers

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/internal/application/services/cart"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
)

type CartHandlers struct {
	logger *logger.Logger
}

func NewCartHandlers(reg *registry.Registry) *CartHandlers {
	return &CartHandlers{
		logger: reg.Inject("logger").(*logger.Logger),
	}
}

func (hs *CartHandlers) HandleCart(context echo.Context) error {
	s := cart.NewService(hs.logger)

	data := new(cart.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		cart.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}
