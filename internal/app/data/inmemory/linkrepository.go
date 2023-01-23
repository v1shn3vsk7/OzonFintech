package inmemory

import (
	"OzonTestTask/internal/app/model"
	"errors"
)

type LinkRepository struct {
	data []Data
}

func (r *LinkRepository) Create(m *model.Link) error {
	if err := m.ValidateURL(); err != nil {
		return err
	}

	if ok := r.checkIfUrlExists(m); ok == true {
		return nil
	}

	data := Data{
		Id:        len(r.data) + 1,
		OriginUrl: m.OriginUrl,
	}
	m.ShortUrl = model.HashUrl(data.Id)
	data.ShortUrl = m.ShortUrl

	r.data = append(r.data, data)

	return nil
}

func (r *LinkRepository) FindByShortURL(m *model.Link) error {
	for i := range r.data {
		if r.data[i].ShortUrl == m.ShortUrl {
			m.OriginUrl = r.data[i].OriginUrl
			return nil
		}
	}
	return errors.New("URL not found")
}

func (r *LinkRepository) checkIfUrlExists(m *model.Link) bool {
	for i := range r.data {
		if r.data[i].OriginUrl == m.OriginUrl {
			m.ShortUrl = r.data[i].ShortUrl
			return true
		}
	}
	return false
}
