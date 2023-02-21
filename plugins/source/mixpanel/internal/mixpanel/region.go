package mixpanel

import (
	"fmt"
	"strings"
)

type Region string

const (
	RegionNone = Region("")
	RegionUS   = Region("us")
	RegionEU   = Region("eu")
)

func ParseRegion(v string) (Region, error) {
	r := Region(strings.ToLower(v))
	if r != RegionNone && r != RegionUS && r != RegionEU {
		return RegionNone, fmt.Errorf("unknown region %q", v)
	}
	return r, nil
}
