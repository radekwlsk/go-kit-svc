package stringsvc

import (
	"context"
	"errors"
	"strings"
	"unicode"
)

// StringService provides operations on strings
type StringService interface {
	TitleCase(context.Context, string) (string, error)
	RemoveWhitespace(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func New() StringService {
	return stringService{}
}

var ErrEmptyString = errors.New("empty string")

func (stringService) TitleCase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.Title(s), nil
}

func (stringService) RemoveWhitespace(_ context.Context, s string) (string, error) {
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

func (stringService) Count(_ context.Context, s string) int {
	return len(s)
}
