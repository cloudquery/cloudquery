package client

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/hashicorp/go-hclog"
)

type retryer struct {
	aws.Retryer
	logger hclog.Logger
}

func newRetryer(logger hclog.Logger, maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return &retryer{
			Retryer: retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = maxRetries
				o.MaxBackoff = time.Second * time.Duration(maxBackoff)
			}),
			logger: logger,
		}
	}
}

func (r *retryer) RetryDelay(attempt int, err error) (time.Duration, error) {
	dur, retErr := r.Retryer.RetryDelay(attempt, err)

	logParams := []interface{}{
		"duration", dur.String(),
		"attempt", attempt,
		"err", err,
	}
	if retErr != nil {
		logParams = append(logParams, "retrier_err", retErr)
		r.logger.Debug("RetryDelay returned err", logParams...)
	} else {
		r.logger.Debug("waiting before retry...", logParams...)
	}
	return dur, retErr
}
