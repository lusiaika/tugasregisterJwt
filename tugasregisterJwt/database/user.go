package database

import (
	"context"
	"database/sql"
	"log"
	"time"
	"tugasregisterjwt/entity"
)

func (s *Database) Register(ctx context.Context, user entity.User) (*entity.UserResult, error) {
	result := &entity.UserResult{}

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := s.SqlDb.QueryContext(ctx, "sp_register",
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("age", user.Age),
		sql.Named("createdat", time.Now()),
		sql.Named("updatedat", time.Now()))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&result.Age,
			&result.Email,
			&result.Id,
			&result.Username,
		)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
