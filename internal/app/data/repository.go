package data

import "OzonTestTask/internal/app/model"

type ILinkRepository interface {
	Create(link *model.Link) error
	FindByShortURL(link *model.Link) error
}
