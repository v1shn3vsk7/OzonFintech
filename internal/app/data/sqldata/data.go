package sqldata

import (
	"OzonTestTask/internal/app/data"
	"database/sql"
	_ "github.com/lib/pq"
)

type Data struct {
	db     		   *sql.DB
	linkRepository *LinkRepository
}

func New(db *sql.DB) *Data {
	return &Data{
		db: db,
	}
}

func (d *Data) Link() data.LinkRepository {
	if d.linkRepository != nil {
		return d.linkRepository
	}

	d.linkRepository = &LinkRepository{
		data: d,
	}

	return d.linkRepository
}
