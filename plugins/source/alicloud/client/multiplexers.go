package client

import (
	"fmt"
	"sort"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

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
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.services {
		region := getRegion(client.services[accountID])
		// Ensure that the region is always set by a region that has been initialized
		if region == "" {
			meta.(*Client).Logger().Trace().Str("accountID", accountID).Msg("no valid regions have been specified for this account")
			continue
		}
		l = append(l, client.WithAccountIDAndRegion(accountID, region))
	}
	return l
}

func ServiceAccountRegionMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for accountID, accountRegions := range client.services {
			regions, ok := ServiceRegions[service]
			if !ok {
				panic(fmt.Sprintf("service %s is not yet supported. Developer hint: add it to client/service_regions.go", service))
			}
			for _, region := range regions {
				if _, regionConfigured := accountRegions[region]; regionConfigured {
					l = append(l, client.WithAccountIDAndRegion(accountID, region))
				}
			}
		}
		return l
	}
}
