package partitions

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type (
	partitions []*partition
	partition  struct {
		Name     string
		Services map[string]*service
	}
	service struct {
		Name    string
		Regions []string
	}
)

func (p partitions) Partitions() []string {
	result := make([]string, 0, len(p))
	for _, pt := range p {
		result = append(result, pt.Name)
	}
	slices.Sort(result)
	return result
}

func (p partitions) Services() []string {
	var services []string
	for _, pt := range p {
		services = append(services, maps.Keys(pt.Services)...)
	}
	slices.Sort(services)
	return slices.Compact(services)
}

func (p partitions) Regions() []string {
	var regions []string
	for _, pt := range p {
		for _, svc := range pt.Services {
			regions = append(regions, svc.Regions...)
		}
	}
	slices.Sort(regions)
	return slices.Compact(regions)
}
