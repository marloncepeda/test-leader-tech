package user

import (
	"api/pkg/errors"
	"api/pkg/querys"
	"database/sql"
	"fmt"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserPortRepositories {
	return &userRepository{db: db}
}

func (r *userRepository) GetUser(id string) ([]UserModel, error) {
	var users []UserModel
	stmt, err := r.db.Prepare(querys.GetUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserModel
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Lastname,
			&user.Username,
			&user.Type,
			&user.Email,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) CreateUser(user UserModel) ([]UserModel, error) {
	var users []UserModel
	stmt, err := r.db.Prepare(querys.PostUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Lastname, user.Username, user.hash, user.Type, user.Email, user.salt)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return nil, err
	}

	id, _ := res.LastInsertId()

	user.Id = int(id)
	users = append(users, user)
	return users, nil
}

func (r *userRepository) UpdateUser(user UserModel) ([]UserModel, error) {
	var users []UserModel
	stmt, err := r.db.Prepare(querys.UpdateUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Lastname, user.Password, user.Username, user.Email, user.Id)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return nil, err
	}

	id, _ := res.LastInsertId()

	user.Id = int(id)
	users = append(users, user)
	return users, nil
}

func (r *userRepository) DeleteUser(id string) (string, error) {
	stmt, err := r.db.Prepare(querys.DeleteUsers)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(id)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return "", err
	}
	fmt.Println(rows)
	return "user deleted", nil
}
