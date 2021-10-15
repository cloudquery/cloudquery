package drift

import (
	"fmt"
	"strings"
)

type Resource struct {
	ID string
}

type ResourceList []*Resource

func (r ResourceList) IDs(exclude ...*Resource) []string {
	exMap := make(map[string]struct{}, len(exclude))
	for i := range exclude {
		exMap[exclude[i].ID] = struct{}{}
	}

	ret := make([]string, 0, len(r))
	for i := range r {
		if _, ok := exMap[r[i].ID]; !ok {
			ret = append(ret, r[i].ID)
		}
	}
	return ret
}

type Result struct {
	Provider     string
	ResourceType string

	Different ResourceList
	DeepEqual ResourceList
	Equal     ResourceList
	Missing   ResourceList
	Extra     ResourceList
}

func (r *Result) String() string {
	stringDump := func(input []*Resource, name string, dst *[]string) {
		switch l := len(input); l {
		case 0:
			return
		case 1:
			*dst = append(*dst, fmt.Sprintf("%d %s (%s)", l, name, input[0].ID))
		case 2:
			*dst = append(*dst, fmt.Sprintf("%d %s (%s, %s)", l, name, input[0].ID, input[1].ID))
		default:
			*dst = append(*dst, fmt.Sprintf("%d %s (%s, %s, ...)", l, name, input[0].ID, input[1].ID))
		}
	}

	var parts []string
	stringDump(r.Different, "different", &parts)
	stringDump(r.Equal, "equal", &parts)
	stringDump(r.DeepEqual, "deepequal", &parts)
	stringDump(r.Missing, "missing", &parts)
	stringDump(r.Extra, "extra", &parts)
	if len(parts) == 0 {
		parts = append(parts, "no")
	}

	return fmt.Sprintf("%s:%s has %s resources", r.Provider, r.ResourceType, strings.Join(parts, ", "))
}

func (r *Result) HasResources() bool {
	return len(r.Different)+len(r.Equal)+len(r.DeepEqual)+len(r.Missing)+len(r.Extra) > 0
}

type Results []*Result

func (rs Results) StringSlice() []string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		if !r.HasResources() {
			continue
		}
		parts = append(parts, r.String())
	}

	if len(parts) == 0 {
		return []string{"no results"}
	}

	return parts
}
