package postgres

import "github.com/jackc/pgx/v5"

type Postgres struct {
	*pgx.Conn
}

func NewPostgres(conn *pgx.Conn) *Postgres {
	return &Postgres{
		Conn: conn,
	}
}
