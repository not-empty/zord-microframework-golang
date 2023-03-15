package http

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/application/services/dummy/GET"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
)

type Server struct {
	Environment string

	config *config.Config
	logger *logger.Logger
}

func NewServer(Environment string) *Server {
	c := pkg.KernelDependencies["config"]
	l := pkg.KernelDependencies["logger"]

	return &Server{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
	}
}

func (hs *Server) Start(port string) {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	server.GET("/v1/dummy/:dummy_id", hs.HandleDummy)
	hs.Shutdown(server.Start(port))

}

func (hs *Server) HandleDummy(context echo.Context) error {
	s := dummy.NewService(hs.logger)
	s.Execute(
		dummy.NewRequest(context.Param("dummy_id")),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
