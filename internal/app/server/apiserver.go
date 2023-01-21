package server

import (
	"OzonTestTask/internal/app/data/sqldata"
	"database/sql"
	"net/http"
)

func Start(cfg *Config) error {
	db, err := newDb(cfg.DbUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	data := sqldata.New(db)
	s := NewServer(data)

	return http.ListenAndServe(cfg.BindAddr, s)
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
