package database

import (
	"context"
	"database/sql"
	"log"
	"tugasregisterjwt/entity"
)

func (s *Database) Login(ctx context.Context, i string) (*entity.User, error) {
	result := &entity.User{}

	rows, err := s.SqlDb.QueryContext(ctx, "select password from users where username = @username",
		sql.Named("username", i))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.Password,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return result, nil
}
