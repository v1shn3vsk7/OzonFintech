package sqldata

import "OzonTestTask/internal/app/model"

type LinkRepository struct {
	data *Data
}

func (r *LinkRepository) Create(m *model.Link) error {
	if err := m.BeforeInsert(); err != nil {
		return err
	}
	//
	r.data.db.QueryRow(
		"INSERT INTO links (origin_link, short_link) VALUES ($1,  $2)",
		m.OriginUrl, m.ShortUrl)

	return nil
}

func (r *LinkRepository) FindByShortLink(link string) (*model.Link, error)  {
	m := model.Link{}

	if err := r.data.db.QueryRow(
		"SELECT id, origin_link, short_link FROM links WHERE short_link = $1",
		link).Scan(
			&m.OriginUrl,
			&m.ShortUrl, //FIX LATER
			);
		 err != nil {
		return nil, err
	}

	return &m, nil
}


