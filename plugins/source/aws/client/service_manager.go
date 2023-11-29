package client

import "github.com/thoas/go-funk"

type ServicesPartitionAccountMap map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services ServicesPartitionAccountMap
}

func (s *ServicesManager) ServicesByPartitionAccount(partition, accountId string) *Services {
	return s.services[partition][accountId]
}

func (s *ServicesManager) InitServices(details svcsDetail) {
	s.InitServicesForPartitionAccount(details.partition, details.accountId, details.svcs)
}

func (s *ServicesManager) InitServicesForPartitionAccount(partition, accountId string, svcs Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = &svcs
	}

	s.services[partition][accountId].Regions = funk.UniqString(append(s.services[partition][accountId].Regions, svcs.Regions...))
}
