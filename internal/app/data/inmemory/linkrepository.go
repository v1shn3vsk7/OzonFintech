package inmemory

import (
	"OzonTestTask/internal/app/model"
	"errors"
)

type LinkRepository struct {
	Data []Data
}

func (r *LinkRepository) Create(m *model.Link) error {
	if err := m.ValidateURL(); err != nil {
		return err
	}

	if ok := r.checkIfUrlExists(m); ok == true {
		return nil
	}

	data := Data{
		Id:        len(r.Data) + 1,
		OriginUrl: m.OriginUrl,
	}
	m.ShortUrl = model.HashUrl(data.Id)
	data.ShortUrl = m.ShortUrl

	r.Data = append(r.Data, data)

	return nil
}

func (r *LinkRepository) FindByShortURL(m *model.Link) error {
	for i := 0; i < len(r.Data); i++ {
		if r.Data[i].ShortUrl == m.ShortUrl {
			m.OriginUrl = r.Data[i].OriginUrl
			return nil
		}
	}

	return errors.New("URL not found")
}

func (r *LinkRepository) checkIfUrlExists(m *model.Link) bool {
	for i := range r.Data {
		if r.Data[i].OriginUrl == m.OriginUrl {
			m.ShortUrl = r.Data[i].ShortUrl
			return true
		}
	}

	return false
}
