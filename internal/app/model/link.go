package model

import (
	"net/url"
)

type Link struct {
	Id        int
	OriginUrl string `json:"URL"`
	ShortUrl  string
}

func HashUrl(id int) string {
	var alp = []rune {'a','b','c','d','e','f','g','h','i',
		'j','k','l','m','n','o','p','q','r','s','t','u','v',
		'w','x','y','z', 'A','B','C','D','E','F','G','H','I',
		'J','K','L','M','N','O','P','Q','R','S','T','U','V',
		'W','X','Y','Z', '0','1','2','3','4','5','6','7','8',
		'9','0', '_'}

	base := len(alp)

	var hash []rune

	for id > 0 {
		hash = append(hash, alp[id % base])
		id /= base
	}

	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}

	return string(hash)
}

func (l *Link) ValidateURL() error {
	_, err := url.ParseRequestURI(l.OriginUrl)
	if err != nil {
		return err
	}

	return nil
}
