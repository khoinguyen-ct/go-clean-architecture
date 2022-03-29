package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/log"
	"net/http"
)

var logger = log.GetLogger()

type Server struct {
	ctx             context.Context
	httpSrv         *http.Server
	server          *echo.Echo
	RouteRegister   *echo.Router
}

func NewHttpServer() *Server {
	var hs Server

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	hs.server = e
	hs.RouteRegister = hs.server.Router()
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%v", config.GetConfiguration().HTTPPort),
	}
	hs.httpSrv = httpServer
	hs.server.Server = httpServer
	hs.registerRoutes()

	return &hs
}

func (sv *Server) Run(ctx context.Context) error {
	sv.ctx = ctx
	if err := sv.server.StartServer(sv.httpSrv); err != nil {
		if err == http.ErrServerClosed {
			logger.Info("server has shutdown gracefully")
		}
		return err
	}
	return nil
}

func (sv *Server) ShutDown(ctx context.Context) {
	if err := sv.server.Shutdown(ctx); err != nil {
		logger.Errorf("failed to shutdown server")
	}
}
