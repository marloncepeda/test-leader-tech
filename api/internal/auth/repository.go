package auth

import (
	"api/pkg/errors"
	"api/pkg/querys"
	"database/sql"
	"log"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthPortRepositories {
	return &authRepository{db: db}
}

func (ar authRepository) GetLogin(auth AuthModel) (AuthModelOut, error) {
	var authMo AuthModelOut
	stmt, err := ar.db.Prepare(querys.GetUsersAuth)
	if err != nil {
		return authMo, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(auth.Username)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return authMo, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&authMo.Id,
			&authMo.Username,
			&authMo.Password,
			&authMo.Salt,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	return authMo, nil
}
