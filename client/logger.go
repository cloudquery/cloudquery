package client

import "github.com/hashicorp/go-hclog"

type awsLogger struct {
	hclog.Logger
	accounts []Account
}

func (a awsLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	for i, arg := range args {
		switch at := arg.(type) {
		case string:
			args[i] = accountObfusactor(a.accounts, at)
		case error:
			args[i] = accountObfusactor(a.accounts, at.Error())
		}
	}
	a.Logger.Log(level, accountObfusactor(a.accounts, msg), args...)
}

func (a *awsLogger) Trace(msg string, args ...interface{}) {
	a.Log(hclog.Trace, msg, args...)
}

func (a *awsLogger) Debug(msg string, args ...interface{}) {
	a.Log(hclog.Debug, msg, args...)
}

func (a *awsLogger) Info(msg string, args ...interface{}) {
	a.Log(hclog.Info, msg, args...)
}

func (a *awsLogger) Warn(msg string, args ...interface{}) {
	a.Log(hclog.Warn, msg, args...)
}

func (a *awsLogger) Error(msg string, args ...interface{}) {
	a.Log(hclog.Error, msg, args...)
}
