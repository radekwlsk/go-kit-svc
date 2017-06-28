package stringsvc

import (
	"errors"
	"strings"
	"unicode"
)

// StringService provides operations on strings
type StringService interface {
	TitleCase(string) (string, error)
	RemoveWhitespace(string) (string, error)
	Count(string) int
}

type stringService struct{}

func New() StringService {
	return stringService{}
}

var ErrEmptyString = errors.New("empty string")

func (stringService) TitleCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.Title(s), nil
}

func (stringService) RemoveWhitespace(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
