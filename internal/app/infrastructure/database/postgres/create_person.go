package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const createPersonSQL = `insert into persons (first_name, last_name, middle_name, phone)
values ($1, $2, $3, $4)
on conflict (first_name, last_name, middle_name, phone)
do update set updated_at=now()
returning id, updated_at, created_at`

func (db *Postgres) CreatePerson(ctx context.Context, person entity.Person) (entity.Person, error) {
	row := db.QueryRow(ctx, createPersonSQL, person.FirstName, person.LastName, person.MiddleName, person.Phone)

	err := row.Scan(&person.ID, &person.UpdatedAt, &person.CreatedAt)
	if err != nil {
		return person, err
	}

	return person, nil
}
