package models

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

type LogEventWrapper struct {
	types.LogEvent
	// An object describing the result of your get relational database log streams request.
	LogStreamName string
}

type DistributionWrapper struct {
	*types.LightsailDistribution
	LatestCacheReset *lightsail.GetDistributionLatestCacheResetOutput
}
