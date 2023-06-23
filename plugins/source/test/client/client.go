package client

import (
	"github.com/rs/zerolog"
)

type TestClient struct {
	Logger zerolog.Logger
}

func (*TestClient) ID() string {
	return "TestClient"
}
