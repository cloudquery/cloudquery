package client

import (
	"sort"

	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

// Extract region from service list
func getRegion(regionalList []string) string {
	// TODO: We should try and find the closest region if possible. This will require checking the following locations:
	// 1. Region defined by the ec2 metadata service
	// 2. Region defined by the AWS_REGION environment variable
	// 3. Region defined by the AWS_DEFAULT_REGION environment variable
	// 4. Region defined by the local config file
	if len(regionalList) == 0 {
		return ""
	}
	regions := append([]string{}, regionalList...)
	// Sorting is important because the plugin SDK requires that multiplexers return a deterministic client. This means if it calls
	// Multiplex multiple times the clients returned should be the same.
	sort.Strings(regions)
	return regions[0]
}

func AccountMultiplex(table string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		l := make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				region := getRegion(client.ServicesManager.services[partition][accountID].Regions)
				// Ensure that the region is always set by a region that has been initialized
				if region == "" {
					// This can only happen if a user specifies a region from a different partition
					meta.(*Client).Logger().Trace().
						Str("accountID", accountID).
						Str("table", table).
						Str("partition", partition).Msg("no valid regions have been specified for this account")
					continue
				}
				l = append(l, client.withPartitionAccountIDAndRegion(partition, accountID, region))
			}
		}
		return l
	}
}

func ServiceAccountRegionMultiplexer(table, service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		notSupportedRegions := make([]string, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for _, region := range client.ServicesManager.services[partition][accountID].Regions {
					if !isSupportedServiceForRegion(service, region) {
						if client.specificRegions {
							notSupportedRegions = append(notSupportedRegions, region)
						}
						client.Logger().Trace().Str("service", service).Str("region", region).Str("table", table).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					l = append(l, client.withPartitionAccountIDAndRegion(partition, accountID, region))
				}
			}
		}
		generateLogMessages(client, table, service, notSupportedRegions, len(l) == 0)
		return l
	}
}

func ServiceAccountRegionsLanguageCodeMultiplex(table, service string, codes []string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		l := make([]schema.ClientMeta, 0)
		accountRegions := ServiceAccountRegionMultiplexer(table, service)(meta)
		for _, c := range accountRegions {
			for _, code := range codes {
				client := c.(*Client).withLanguageCode(code)
				l = append(l, client)
			}
		}
		return l
	}
}

func ServiceAccountRegionNamespaceMultiplexer(table, service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		notSupportedRegions := make([]string, 0)
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				for _, region := range client.ServicesManager.services[partition][accountID].Regions {
					if !isSupportedServiceForRegion(service, region) {
						if client.specificRegions {
							notSupportedRegions = append(notSupportedRegions, region)
						}
						client.Logger().Trace().Str("service", service).Str("region", region).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					for _, ns := range AllNamespaces {
						l = append(l, client.withPartitionAccountIDRegionAndNamespace(partition, accountID, region, ns))
					}
				}
			}
		}
		generateLogMessages(client, table, service, notSupportedRegions, len(l) == 0)
		return l
	}
}

func ServiceAccountRegionScopeMultiplexer(table, service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		notSupportedRegions := make([]string, 0)
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for partition := range client.ServicesManager.services {
			for accountID := range client.ServicesManager.services[partition] {
				// always fetch cloudfront related resources
				switch partition {
				case "aws":
					l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, awsCloudfrontScopeRegion, wafv2types.ScopeCloudfront))
				case "aws-cn":
					l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, awsCnCloudfrontScopeRegion, wafv2types.ScopeCloudfront))
				}

				for _, region := range client.ServicesManager.services[partition][accountID].Regions {
					if !isSupportedServiceForRegion(service, region) {
						if client.specificRegions {
							notSupportedRegions = append(notSupportedRegions, region)
						}
						client.Logger().Trace().Str("service", service).Str("region", region).Str("partition", partition).Msg("region is not supported for service")
						continue
					}
					l = append(l, client.withPartitionAccountIDRegionAndScope(partition, accountID, region, wafv2types.ScopeRegional))
				}
			}
		}
		generateLogMessages(client, table, service, notSupportedRegions, len(l) == 0)
		return l
	}
}

func generateLogMessages(client *Client, table, service string, skippedRegions []string, emptyMultiplexer bool) {
	if len(skippedRegions) == 0 {
		return
	}
	loggerEvent := client.Logger().Info()
	if emptyMultiplexer {
		loggerEvent = client.Logger().Warn()
	}
	loggerEvent.Str("service", service).
		Str("table", table).
		Strs("skipped regions", skippedRegions).
		Strs("supported regions", supportedRegions(service)).
		Msg("skipping table for unsupported regions. To fix this message, ensure to configure only supported regions for the table")
}
