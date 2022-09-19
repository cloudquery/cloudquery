package client

import "context"

type NoRateLimiter struct {
}

func (*NoRateLimiter) GetToken(ctx context.Context, cost uint) (func() error, error) {
	return func() error {
		return nil
	}, nil
}

func (*NoRateLimiter) AddTokens(uint) error {
	return nil
}
