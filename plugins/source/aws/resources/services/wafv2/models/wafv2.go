package models

import "github.com/aws/aws-sdk-go-v2/service/wafv2/types"

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}
