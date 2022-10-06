package models

import "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"

type CloudTrailWrapper struct {
	types.Trail
	Tags map[string]string
}
