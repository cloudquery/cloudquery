package models

import "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"

type ELBv1LoadBalancerWrapper struct {
	types.LoadBalancerDescription
	Tags       map[string]string
	Attributes *types.LoadBalancerAttributes
}
