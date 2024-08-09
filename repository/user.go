package repository

import (
	"database/sql"
	"golang_project/entity"
	"golang_project/repository/models"
)

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

func (s *userStorage) New(newPerson *entity.Person) (int64, error) {
	// Begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return 0, ErrCouldNotBeginTx
	}
	defer tx.Rollback() // Rollback if any error occurs

	// Insert into person table
	personStmt := `INSERT INTO person (name) VALUES ($1) RETURNING id`
	var personID int64
	err = tx.QueryRow(personStmt, newPerson.Name).Scan(&personID)
	if err != nil {
		return 0, ErrCouldNotInsertPerson
	}

	// Insert into phone table
	phoneStmt := `INSERT INTO phone (number, person_id) VALUES ($1, $2)`
	_, err = tx.Exec(phoneStmt, newPerson.PhoneNumber, personID)
	if err != nil {
		return 0, ErrCouldNotInsertPhone
	}

	// Insert into address table
	addressStmt := `INSERT INTO address (city, state, street1, street2, zip_code) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var addressID int64
	err = tx.QueryRow(addressStmt, newPerson.City, newPerson.State, newPerson.Street1, newPerson.Street2, newPerson.ZipCode).Scan(&addressID)
	if err != nil {
		return 0, ErrCouldNotInsertAddress
	}

	// Insert into address_join table
	addressJoinStmt := `INSERT INTO address_join (person_id, address_id) VALUES ($1, $2)`
	_, err = tx.Exec(addressJoinStmt, personID, addressID)
	if err != nil {
		return 0, ErrCouldNotInsertAddressJoin
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return 0, ErrCouldNotCommitTx
	}

	return personID, nil
}

func (s *userStorage) GetPerson(personID int64) (*entity.Person, error) {
	stmt := `SELECT
				p.name AS Name,
				ph.number AS PhoneNumber,
				a.city AS City,
				a.state AS State,
				a.street1 AS Street1,
				a.street2 AS Street2,
				a.zip_code AS ZipCode
			FROM
				person p
			JOIN
				phone ph ON p.id = ph.person_id
			JOIN
				address_join aj ON p.id = aj.person_id
			JOIN
				address a ON aj.address_id = a.id
			WHERE
				p.id = $1`

	query, err := s.db.Prepare(stmt)
	if err != nil {
		return nil, ErrCouldNotPrepareStmt
	}
	defer query.Close()

	var queryResult models.Person
	err = query.QueryRow(personID).Scan(
		&queryResult.Name,
		&queryResult.PhoneNumber,
		&queryResult.City, &queryResult.State, &queryResult.Street1,
		&queryResult.Street2, &queryResult.ZipCode,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, ErrCouldNotQueryScanData
	}
	personObj := entity.Person{
		Id:          queryResult.ID,
		Name:        queryResult.Name,
		PhoneNumber: queryResult.PhoneNumber,
		City:        queryResult.City,
		State:       queryResult.State,
		Street1:     queryResult.Street1,
		Street2:     queryResult.Street2,
		ZipCode:     queryResult.ZipCode,
		CreatedAt:   queryResult.CreatedAt,
		UpdatedAt:   queryResult.UpdatedAt.Time,
	}

	return &personObj, nil
}
