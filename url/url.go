package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
}

func Parse(rawurl string) (*URL, error) {
	i := strings.Index(rawurl, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme := rawurl[:i]
	return &URL{scheme}, nil
}
