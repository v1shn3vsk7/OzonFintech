package server

import (
	"OzonTestTask/internal/app/data/sqldata"
	"database/sql"
	"errors"
	"net/http"
	"os"
)

func Start(cfg *Config) error {
	storeType := os.Getenv("STORE_TYPE")
	if storeType == "inmemory" {
		data := cfg.Data

		s := NewServer(data)

		return http.ListenAndServe(cfg.BindAddr, s)

	} else if storeType == "postgres" {
		db, err := newDb(cfg.DbUrl)
		if err != nil {
			return err
		}

		defer db.Close()

		data := sqldata.New(db)
		s := NewServer(data)

		return http.ListenAndServe(cfg.BindAddr, s)
	} else {
		return errors.New("no choice for storage type")
	}

}

func newDb(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
