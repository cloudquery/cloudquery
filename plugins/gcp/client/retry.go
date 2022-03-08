package client

import (
	"context"
	"errors"
	"reflect"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/googleapis/gax-go/v2"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/googleapi"
)

func shouldRetryFunc(log hclog.Logger) func(err error) bool {
	return func(err error) bool {
		if IgnoreErrorHandler(err) {
			reason := ""
			var gerr *googleapi.Error
			if errors.As(err, &gerr) && len(gerr.Errors) > 0 {
				reason = gerr.Errors[0].Reason
			}

			log.Debug("retrier not retrying: ignore error", "err", err, "err_reason", reason)
			return false
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Debug("retrier not retrying", "err", err)
			return false
		}

		log.Debug("retrying error", "err", err)
		return true
	}
}

// RetryingResolver runs the TableResolver with retry. Not very good as it could cause multiple resources with multi-page retries (retrying after fetching some resources)
func RetryingResolver(f schema.TableResolver) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		cl := meta.(*Client)
		return gax.Invoke(ctx, func(ctx context.Context, _ gax.CallSettings) error {
			return f(ctx, meta, parent, res)
		}, gax.WithRetry(func() gax.Retryer {
			return gax.OnErrorFunc(cl.backoff.Gax, shouldRetryFunc(cl.logger))
		}))
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
		return gax.OnErrorFunc(c.backoff.Gax, shouldRetryFunc(c.logger))
	}))
	return val, err
}

type doer interface {
	Do(...googleapi.CallOption) (interface{}, error)
	Context(context.Context) interface{}
}

func makeDoer(x interface{}) doer {
	return &doerWrapper{doer: x}
}

type doerWrapper struct {
	doer interface{}
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
