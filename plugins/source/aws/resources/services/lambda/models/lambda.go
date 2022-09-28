package models

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type AliasWrapper struct {
	*types.AliasConfiguration
	UrlConfig *lambda.GetFunctionUrlConfigOutput
}

type RuntimeWrapper struct {
	Name string
}
