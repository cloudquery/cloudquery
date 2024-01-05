package client

import (
	"sync"

	"github.com/thoas/go-funk"
)

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	m         sync.Map // implies usage by reference
	hasValues bool
}

func (s *ServicesManager) ServicesByPartitionAccount(partition, accountID string) *Services {
	p, ok := s.m.Load(partition)
	if !ok {
		return nil
	}

	svc, ok := p.(*sync.Map).Load(accountID)
	if !ok {
		return nil
	}

	return svc.(*Services)
}

func (s *ServicesManager) InitServices(details svcsDetail) {
	s.InitServicesForPartitionAccount(details.partition, details.accountId, details.svcs)
}

func (s *ServicesManager) InitServicesForPartitionAccount(partition, accountID string, svcs Services) {
	p, _ := s.m.LoadOrStore(partition, new(sync.Map))
	svc, loaded := p.(*sync.Map).LoadOrStore(accountID, &svcs)
	if loaded {
		svc := svc.(*Services)
		svc.Regions = funk.UniqString(append(svc.Regions, svcs.Regions...))
	}
	s.hasValues = true
}

func (s *ServicesManager) Range(f func(partition, accountID string, svc *Services)) {
	s.m.Range(func(p, pMap any) bool {
		partition := p.(string)
		pMap.(*sync.Map).Range(func(a, s any) bool {
			accountID := a.(string)
			svc := s.(*Services)
			f(partition, accountID, svc)
			return true
		})
		return true
	})
}
