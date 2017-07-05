package stringsvc

import (
	"time"

	"context"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

// NewLoggingMiddleware returns StringService middleware that logs
// information about each method execution including:
// method name, input, output, error if present and time of execution
func NewLoggingMiddleware(svc StringService, logger log.Logger) StringService {
	return loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) TitleCase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "title_case",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.TitleCase(ctx, s)
	return
}

func (mw loggingMiddleware) RemoveWhitespace(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "remove_whitespace",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.RemoveWhitespace(ctx, s)
	return
}

func (mw loggingMiddleware) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(ctx, s)
	return
}
