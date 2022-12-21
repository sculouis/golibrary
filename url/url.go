package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawurl string) (*URL, error) {
	i := strings.Index(rawurl, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme, rest := rawurl[:i], rawurl[i+3:]
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i > 0 {
		host = rest[:i]
		path = rest[i+1:]
	}
	return &URL{scheme, host, path}, nil
}
