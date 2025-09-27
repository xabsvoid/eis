package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/service"
	"github.com/xabsvoid/eis/internal/app/infrastructure/database/postgres"
	"github.com/xabsvoid/eis/internal/app/infrastructure/transport/http"
)

const (
	EnvDSN  = "DSN"
	EnvHost = "HOST"
)

func main() {
	flagDSN := flag.String("dsn", "postgres://user:pwd@host:5432/db", "dsn db")
	flagHost := flag.String("host", ":8080", "server host")
	flag.Parse()

	dsn := os.Getenv(EnvDSN)
	if flagDSN != nil {
		dsn = *flagDSN
	}

	host := os.Getenv(EnvHost)
	if flagHost != nil {
		host = *flagHost
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	repository := postgres.NewPostgres(conn)

	appService := service.NewService(repository)

	httpHandlers := http.NewHandlers(appService)

	httpServer := echo.New()

	http.RegisterHandlers(httpServer, httpHandlers)

	err = httpServer.Start(host)
	if err != nil {
		log.Fatal(err)
	}
}
