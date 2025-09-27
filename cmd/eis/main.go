package main

import (
	"context"
	"flag"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/service"
	"github.com/xabsvoid/eis/internal/app/infrastructure/database/postgres"
	"github.com/xabsvoid/eis/internal/app/infrastructure/transport/http"
)

func main() {
	dsn := flag.String("dsn", "postgres://user:pwd@host:5432/db", "dsn db")
	host := flag.String("host", ":8080", "server host")
	flag.Parse()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, *dsn)
	if err != nil {
		log.Fatal(err)
	}

	repository := postgres.NewPostgres(conn)

	appService := service.NewService(repository)

	httpHandlers := http.NewHandlers(appService)

	httpServer := echo.New()

	http.RegisterHandlers(httpServer, httpHandlers)

	err = httpServer.Start(*host)
	if err != nil {
		log.Fatal(err)
	}
}
