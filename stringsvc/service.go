package stringsvc

// Service interface definition and basic service methods implementation,
// the actual actions performed by service on data.

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

// New returns a stateless implementation of StringService
func New() StringService {
	return stringService{}
}

// ErrEmptyString protects the TitleCase and RemoveWhitespace
// from executing on empty strings that will have no effect
var ErrEmptyString = errors.New("empty string")

// TitleCase implements StringService
func (stringService) TitleCase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.Title(s), nil
}

// RemoveWhitespace implements StringService
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

// Count implements StringService
func (stringService) Count(_ context.Context, s string) int {
	return len(s)
}
