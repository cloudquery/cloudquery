package models

import "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"

type ConfigurationOptionDescriptionWrapper struct {
	types.ConfigurationOptionDescription
	ApplicationArn string
}

type ConfigurationSettingsDescriptionWrapper struct {
	types.ConfigurationSettingsDescription
	ApplicationArn string
}
