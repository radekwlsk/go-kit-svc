package stringsvc

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func NewLoggingMiddleware(svc StringService, logger log.Logger) StringService {
	return loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) TitleCase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "title_case",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.TitleCase(s)
	return
}

func (mw loggingMiddleware) RemoveWhitespace(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "remove_whitespace",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.RemoveWhitespace(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
