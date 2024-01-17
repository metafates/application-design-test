package app

import (
	"github.com/gorilla/mux"
	"github.com/metafates/application-design-test/internal/controller/http/middleware"
	v1 "github.com/metafates/application-design-test/internal/controller/http/v1"
	"github.com/metafates/application-design-test/internal/usecase/repo/memmap"
	"github.com/metafates/application-design-test/internal/usecase/service"
	"github.com/metafates/application-design-test/pkg/httpserver"
	"log/slog"
	"os"
	"os/signal"
)

const httpServerPort = "1234"

func Run() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}),
	)

	repo := memmap.New(logger)
	hotelBookingService := service.NewHotelBooking(logger, repo)

	router := mux.NewRouter()
	router.Use(middleware.Logging(logger))

	v1.RegisterRoutes(
		router.PathPrefix("/v1").Subrouter(),
		hotelBookingService,
	)

	httpServer := httpserver.New(
		router,
		httpserver.WithPort(httpServerPort),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go httpServer.Start()
	logger.Info("http server is running", "port", httpServerPort)

	select {
	case <-interrupt:
		logger.Info("received interrupt, shutting down")
	case err := <-httpServer.Notify():
		logger.Error("http server", "err", err)
	}

	if err := httpServer.Shutdown(); err != nil {
		logger.Error("http server shutdown", "err", err)
	}
}
