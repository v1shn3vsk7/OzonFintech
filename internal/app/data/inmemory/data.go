package inmemory

import "OzonTestTask/internal/app/data"

type Data struct {
	Id             int
	OriginUrl      string
	ShortUrl       string
	linkRepository *LinkRepository
}

func (d *Data) Link() data.LinkRepository {
	if d.linkRepository != nil {
		return d.linkRepository
	}

	d.linkRepository = &LinkRepository{
	}

	return d.linkRepository
}

