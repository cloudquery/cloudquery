package main

import (
	"github.com/cloudquery/cloudquery/cmd"
	"github.com/getsentry/sentry-go"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			sentry.CurrentHub().Recover(err)
			panic(err)
		}
	}()
	cmd.Execute()
}
