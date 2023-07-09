package shipment

import (
	"api/pkg/errors"
	"api/pkg/querys"
	"database/sql"
	"log"
)

type shipmentRepository struct {
	db *sql.DB
}

func NewShipmentRepository(db *sql.DB) ShipmentPortRepositories {
	return &shipmentRepository{db: db}
}

func (r *shipmentRepository) GetShipment(id string) ([]ShipmentModel, error) {
	var users []ShipmentModel
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
		var user ShipmentModel
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Lastname,
			&user.Username,
			&user.Email,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *shipmentRepository) CreateShipment(user ShipmentModel) ([]ShipmentModel, error) {
	var users []ShipmentModel
	stmt, err := r.db.Prepare(querys.PostUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Lastname, user.hash, user.Username, user.Email, user.salt)
	if err != nil {
		log.Printf(errors.QueryError.Error(), err.Error())
		return nil, err
	}

	id, _ := res.LastInsertId()

	user.Id = int(id)
	users = append(users, user)
	return users, nil
}

func (r *shipmentRepository) UpdateShipment(user ShipmentModel) ([]ShipmentModel, error) {
	var users []ShipmentModel
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
