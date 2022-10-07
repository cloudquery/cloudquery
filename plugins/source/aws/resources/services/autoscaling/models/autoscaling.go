package models

import "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"

type AutoScalingGroupWrapper struct {
	types.AutoScalingGroup
	NotificationConfigurations []types.NotificationConfiguration
}
