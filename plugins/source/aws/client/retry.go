package client

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/rs/zerolog"
)

type retryer struct {
	aws.Retryer
	logger zerolog.Logger
}

func newRetryer(logger zerolog.Logger, maxRetries int, maxBackoff int) func() aws.Retryer {
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
		r.logger.Debug().Err(retErr).Interface("logParams", logParams).Msg("RetryDelay returned err")
	} else {
		r.logger.Debug().Interface("logParams", logParams).Msg("waiting before retry...")
	}
	return dur, retErr
}
