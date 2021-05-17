package resources

import "github.com/aws/aws-sdk-go-v2/service/route53/types"

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]interface{}
}

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags            map[string]interface{}
	DelegationSetId *string
	VPCs            []types.VPC
}

func getRoute53tagsByResourceID(id string, set []types.ResourceTagSet) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.Tags
		}
	}
	return nil
}
