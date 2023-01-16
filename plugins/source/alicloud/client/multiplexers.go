package client

import (
	"sort"

	"github.com/cloudquery/plugin-sdk/schema"
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
		regionID := getRegion(client.services[accountID])
		// Ensure that the region is always set by a region that has been initialized
		if regionID == "" {
			meta.(*Client).Logger().Trace().Str("accountID", accountID).Msg("no valid regions have been specified for this account")
			continue
		}
		l = append(l, client.withAccountIDAndRegion(accountID, regionID))
	}
	return l
}

func AccountRegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID, account := range client.services {
		for regionID := range account {
			l = append(l, client.withAccountIDAndRegion(accountID, regionID))
		}
	}
	return l
}
