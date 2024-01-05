package client

import "github.com/thoas/go-funk"

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager map[string]map[string]*Services

func (s ServicesManager) ServicesByPartitionAccount(partition, accountId string) *Services {
	return s[partition][accountId]
}

func (s ServicesManager) InitServices(details svcsDetail) {
	s.InitServicesForPartitionAccount(details.partition, details.accountId, details.svcs)
}

func (s ServicesManager) InitServicesForPartitionAccount(partition, accountID string, svcs *Services) {
	p, ok := s[partition]
	if !ok {
		p = make(map[string]*Services)
		s[partition] = p
	}

	svc, ok := p[accountID]
	if !ok {
		p[accountID] = svcs
		return
	}

	svc.Regions = funk.UniqString(append(svc.Regions, svcs.Regions...))
}
