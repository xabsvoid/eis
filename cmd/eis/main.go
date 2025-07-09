package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/service"
	"github.com/xabsvoid/eis/internal/app/infrastructure/database/inmem"
	"github.com/xabsvoid/eis/internal/app/infrastructure/transport/http"
)

func main() {
	host := flag.String("host", ":8080", "server host")
	flag.Parse()

	inMemRepository := inmem.NewInMem()

	appService := service.NewService(inMemRepository)

	httpHandlers := http.NewHandlers(appService)

	httpServer := echo.New()

	http.RegisterHandlers(httpServer, httpHandlers)

	err := httpServer.Start(*host)
	if err != nil {
		log.Fatal(err)
	}
}
