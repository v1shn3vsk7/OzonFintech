package model

import "net/url"

type Link struct {
	OriginUrl string `json:"URL"`
	ShortUrl  string `json:"sURL"`
}

func (l *Link) BeforeInsert() error {
	if err := validateUrl(l.OriginUrl); err != nil {
		return err
	}

	l.ShortUrl = hashLink(l.OriginUrl)

	return nil
}

// HashLink TODO implement creating of short link
func hashLink(link string) string {

	return "hui"
}

func validateUrl(Url string) error {
	_, err := url.ParseRequestURI(Url)
	if err != nil {
		return err
	}

	return nil
}
