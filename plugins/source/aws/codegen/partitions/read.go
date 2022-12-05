package partitions

import (
	"embed"
	"encoding/json"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type (
	jsonRoot struct {
		Partitions map[string]jsonPartition `json:"partitions"`
	}
	jsonPartition struct {
		Services map[string]jsonService `json:"services"`
	}
	jsonService struct {
		Regions map[string]json.RawMessage `json:"regions"`
	}
)

const (
	partitionServiceRegionFile = "data/partition_service_region.json"
)

var (
	//go:embed data/partition_service_region.json
	partitionsFile embed.FS
)

func readPartitions() (partitions, error) {
	f, err := partitionsFile.Open(partitionServiceRegionFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	root := new(jsonRoot)
	err = json.NewDecoder(f).Decode(root)
	if err != nil {
		return nil, err
	}

	var result partitions
	for name, part := range root.Partitions {
		pt := &partition{
			Name:     name,
			Services: make(map[string]*service),
		}
		for svcName, s := range part.Services {
			regions := maps.Keys(s.Regions)
			slices.Sort(regions)

			pt.Services[svcName] = &service{
				Name:    svcName,
				Regions: regions,
			}
		}
		result = append(result, pt)
	}
	return result, nil
}
