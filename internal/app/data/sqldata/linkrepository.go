package sqldata

import (
	"OzonTestTask/internal/app/model"
)

type LinkRepository struct {
	data *Data
}

//разобраться с дупликатами
func (r *LinkRepository) Create(m *model.Link) error {
	if err := m.ValidateURL(); err != nil {
		return err
	}

	if err := r.data.db.QueryRow(
		"INSERT INTO links (origin_link) VALUES ($1) RETURNING id",
		m.OriginUrl).Scan(&m.Id); err != nil {
		return err
	}

	m.ShortUrl = model.HashUrl(m.Id)

	r.data.db.QueryRow(
		"UPDATE links SET short_link = $1 WHERE origin_link = $2",
		m.ShortUrl, m.OriginUrl)

	return nil
}

func (r *LinkRepository) FindByShortURL(m *model.Link) error {
	if err := r.data.db.QueryRow(
		"SELECT origin_link FROM links WHERE short_link = $1",
		m.ShortUrl).Scan(
			&m.OriginUrl);
		 err != nil {
		return err
	}

	return nil
}


