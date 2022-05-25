package client

import (
	"context"
	"errors"
	"net/http"
	"reflect"

	"github.com/googleapis/gax-go/v2"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/googleapi"
)

type doer interface {
	Do(...googleapi.CallOption) (interface{}, error)
	Context(context.Context) interface{}
}

type doerWrapper struct {
	doer interface{}
}

func shouldRetryFunc(log hclog.Logger, maxRetries int) func(err error) bool {
	totalRetries := 0
	return func(err error) bool {
		totalRetries++
		if totalRetries > maxRetries {
			log.Debug("retrier not retrying, reached max retries", "err", err, "max_retries", maxRetries)
			return false
		}
		var gerr *googleapi.Error
		if ok := errors.As(err, &gerr); ok {
			if gerr.Code == http.StatusForbidden {
				var reason string
				if len(gerr.Errors) > 0 {
					reason = gerr.Errors[0].Reason
				}
				log.Debug("retrier not retrying: ignore error", "err", err, "err_reason", reason, "total_retries", totalRetries, "max_retries", maxRetries)
				return false
			}
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Debug("retrier not retrying", "err", err, "total_retries", totalRetries, "max_retries", maxRetries)
			return false
		}

		log.Debug("retrying api call", "err", err, "total_retries", totalRetries, "max_retries", maxRetries)
		return true
	}
}

// RetryingDo runs the given doerIface with retry
// doerIface needs to have two methods: `Call(...googleapi.CallOption) (T, error)` and `Context(ctx.Context) T`
func (c *Client) RetryingDo(ctx context.Context, doerIface interface{}, opts ...googleapi.CallOption) (interface{}, error) {
	var val interface{}
	err := gax.Invoke(ctx, func(ctx context.Context, _ gax.CallSettings) error {
		doer := makeDoer(doerIface)

		var err error
		_ = doer.Context(ctx)
		val, err = doer.Do(opts...)
		return err
	}, gax.WithRetry(func() gax.Retryer {
		return gax.OnErrorFunc(c.backoff.Gax, shouldRetryFunc(c.logger, c.backoff.MaxRetries))
	}))
	return val, err
}

func makeDoer(x interface{}) doer {
	return &doerWrapper{doer: x}
}

func (d *doerWrapper) Do(opts ...googleapi.CallOption) (interface{}, error) {
	do, ok := reflect.TypeOf(d.doer).MethodByName("Do")
	if !ok {
		panic("passed struct doesn't have a Do method")
	}
	if do.Type.NumOut() != 2 {
		panic("passed struct's Do method doesn't return 2 values")
	}

	ret := do.Func.CallSlice([]reflect.Value{
		reflect.ValueOf(d.doer),
		reflect.ValueOf(opts),
	})
	if ret[1].IsNil() {
		return ret[0].Interface(), nil
	}

	return ret[0].Interface(), ret[1].Interface().(error)
}

func (d *doerWrapper) Context(ctx context.Context) interface{} {
	c, ok := reflect.TypeOf(d.doer).MethodByName("Context")
	if !ok {
		panic("passed struct doesn't have a Context method")
	}
	if c.Type.NumOut() != 1 {
		panic("passed struct's Context method doesn't return 1 value")
	}

	ret := c.Func.Call([]reflect.Value{
		reflect.ValueOf(d.doer),
		reflect.ValueOf(ctx),
	})
	return ret[0].Interface()
}
