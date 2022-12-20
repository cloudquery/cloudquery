package client

import (
	"sort"

	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

// Extract region from service list
func getRegion(regionalMap map[string]*Services) string {
	if len(regionalMap) == 0 {
		return ""
	}
	regions := make([]string, 0)
	for i := range regionalMap {
		regions = append(regions, i)
	}
	sort.Strings(regions)
	return regions[0]
}

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	l := make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for partition := range client.ServicesManager.services {
		for accountID := range client.ServicesManager.services[partition] {
			region := getRegion(client.ServicesManager.services[partition][accountID])
			// Ensure that the region is always set by a region that has been initialized
			if region == "" {
				meta.(*Client).Logger().Trace().Str("accountID", accountID).Str("partition", partition).Msg("no valid regions have been specified for this account")
				continue
			}
			l = append(l, client.withPartitionAccountIDAndRegion(partition, accountID, region))
		}
	}
	return l
}

func ServiceAccountRegionMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.(*Client).Logger().Trace().Str("service", service).Str("region", region).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					l = append(l, client.withPartitionAccountIDAndRegion(partition, accountID, region))
				}
			}
		}
		return l
	}
}

func ServiceAccountRegionNamespaceMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.(*Client).Logger().Trace().Str("service", service).Str("region", region).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					for _, ns := range AllNamespaces {
						l = append(l, client.withPartitionAccountIDRegionAndNamespace(partition, accountID, region, ns))
					}
				}
			}
		}
		return l
	}
}

func ServiceAccountRegionScopeMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				// always fetch cloudfront related resources
				l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, cloudfrontScopeRegion, wafv2types.ScopeCloudfront))
				for region := range client.ServicesManager.services[partition][accountID] {
					if !isSupportedServiceForRegion(service, region) {
						meta.(*Client).Logger().Trace().Str("service", service).Str("region", region).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, region, wafv2types.ScopeRegional))
				}
			}
		}
		return l
	}
}
