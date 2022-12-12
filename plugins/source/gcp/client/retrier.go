package client

import (
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Retrier struct {
	backoff    gax.Backoff
	curRetrier int
	maxRetries int
	codes      []codes.Code
}

func (r *Retrier) Retry(err error) (time.Duration, bool) {
	st, ok := status.FromError(err)
	if !ok {
		return 0, false
	}
	c := st.Code()
	for _, rc := range r.codes {
		if c == rc {
			if r.curRetrier >= r.maxRetries {
				return 0, false
			}
			r.curRetrier++
			return r.backoff.Pause(), true
		}
	}
	return 0, false
}
