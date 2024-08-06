package repository

import (
	"database/sql"
	"golang_project/config"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Storage struct {
	db  *sql.DB
	cfg config.Config
}

var (
	ErrCouldNotCreateStmt    = errors.New("could not create statement")
	ErrCouldNotPrepareStmt   = errors.New("could not prepare statement")
	ErrCouldNotQueryResource = errors.New("could not query resource")
	ErrCouldFetchLastInsert  = errors.New("could not fetch last insert")
	ErrCouldNotScanData      = errors.New("could not scan data")
	ErrCouldNotQueryScanData = errors.New("could not query or scan data")
	ErrCouldNotInsert        = errors.New("could not insert data")
	ErrCouldNotUpdate        = errors.New("could not update data")
	ErrCouldNotDelete        = errors.New("could not delete data")
	ErrDuplicateEntry        = errors.New("that record already exist")
)

var (
	ErrCouldNotBeginTx           = errors.New("could not begin transaction")
	ErrCouldNotInsertPerson      = errors.New("could not insert into person table")
	ErrCouldNotInsertPhone       = errors.New("could not insert into phone table")
	ErrCouldNotInsertAddress     = errors.New("could not insert into address table")
	ErrCouldNotInsertAddressJoin = errors.New("could not insert into address_join table")
	ErrCouldNotCommitTx          = errors.New("could not commit transaction")
)

func NewStorage(db *sql.DB, cfg config.Config) (*Storage, error) {
	if pingErr := db.Ping(); pingErr != nil {
		return nil, errors.WithStack(pingErr)
	}
	return &Storage{
		db:  db,
		cfg: cfg,
	}, nil
}

func (s *Storage) CloseDB() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}
