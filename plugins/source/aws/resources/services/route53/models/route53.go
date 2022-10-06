package models

import "github.com/aws/aws-sdk-go-v2/service/route53/types"

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags            map[string]string
	DelegationSetId *string
	VPCs            []types.VPC
}
