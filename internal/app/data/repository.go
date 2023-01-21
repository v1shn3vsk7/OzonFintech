package data

import "OzonTestTask/internal/app/model"

type LinkRepository interface {
	Create(link *model.Link) error
	FindByShortURL(link *model.Link) error
}
