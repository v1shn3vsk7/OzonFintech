package server

import (
	"OzonTestTask/internal/app/data/inmemory"
	"OzonTestTask/internal/app/data/sqldata"
	"database/sql"
	"errors"
	"net/http"
)

func Start(cfg *Config, dataType string) error {
	if dataType == "inmemory" {
		data := &inmemory.Data{}

		s := NewServer(data)

		return http.ListenAndServe(cfg.BindAddr, s)

	} else if dataType == "postgres" {
		db, err := newDb(cfg.DbUrl)
		if err != nil {
			return err
		}

		defer db.Close()

		data := sqldata.New(db)
		s := NewServer(data)

		return http.ListenAndServe(cfg.BindAddr, s)
	} else {
		return errors.New("no choice for data")
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
