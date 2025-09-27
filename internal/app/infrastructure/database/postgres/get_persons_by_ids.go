package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const getPersonsByIDsSQL = `select id, first_name, last_name, middle_name, phone, updated_at, created_at
from persons
where id=ANY($1)`

func (db *Postgres) GetPersonsByIDs(ctx context.Context, ids []int64) ([]entity.Person, error) {
	var persons []entity.Person

	rows, err := db.Query(ctx, getPersonsByIDsSQL, ids)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var person entity.Person
		err = rows.Scan(
			&person.ID,
			&person.FirstName,
			&person.LastName,
			&person.MiddleName,
			&person.Phone,
			&person.UpdatedAt,
			&person.CreatedAt)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return persons, nil
}
