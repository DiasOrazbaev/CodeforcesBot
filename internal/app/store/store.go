package store

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Config *Config
	Db     *sqlx.DB
}

func NewStore(config *Config) *Store {
	return &Store{Config: config}
}

func (s *Store) Open() error {
	dbr, _ := sql.Open("sqlite3", ":memory:")
	db := sqlx.NewDb(dbr, "sqlite3")
	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db
	return nil
}

func (s *Store) Close() error {
	return s.Db.Close()
}
