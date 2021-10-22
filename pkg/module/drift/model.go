package drift

import (
	"fmt"
	"strings"
)

type RunParams struct {
	Debug bool

	TfBackendName, TfMode, TfProvider string
	ForceDeep                         bool
}

type Resource struct {
	ID string `json:"id"`
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
	IAC          string `json:"iac"`
	Provider     string `json:"provider"`
	ResourceType string `json:"resource_type"`

	// Deep mode
	Different ResourceList `json:"diff,omitempty"`       // Resources don't match fully (id + attributes don't match)
	DeepEqual ResourceList `json:"deep_equal,omitempty"` // Resource exists in both places (attributes match)

	// Shallow mode
	Equal ResourceList `json:"equal,omitempty"` // Resource exists in both places (attributes not checked)

	// Both modes
	Missing ResourceList `json:"missing"` // Missing in cloud provider, defined in iac
	Extra   ResourceList `json:"extra"`   // Exists in cloud provider, not defined in iac
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

type Results []*Result

func (rs Results) String() string {
	type combined struct {
		IAC          string
		Provider     string
		ResourceType string
		ResourceIDs  []string
	}
	var combo struct {
		Different []combined
		Extra     []combined
		Equal     []combined
		DeepEqual []combined
		Missing   []combined
	}
	transform := func(r *Result, l ResourceList, dst *[]combined) {
		ids := l.IDs()
		if len(ids) == 0 {
			return
		}
		*dst = append(*dst, combined{
			IAC:          r.IAC,
			ResourceType: r.ResourceType,
			Provider:     r.Provider,
			ResourceIDs:  ids,
		})
	}

	for _, r := range rs {
		if r == nil {
			continue
		}
		transform(r, r.Different, &combo.Different)
		transform(r, r.Extra, &combo.Extra)
		transform(r, r.Equal, &combo.Equal)
		transform(r, r.DeepEqual, &combo.DeepEqual)
		transform(r, r.Missing, &combo.Missing)
	}

	var (
		lines   []string
		summary []string
		total   int
	)

	for _, data := range []struct {
		title string
		list  []combined
	}{
		{
			"not managed by $iac",
			combo.Extra,
		},
		{
			"found in $iac state but missing on the cloud provider",
			combo.Missing,
		},
		{
			"managed by $iac but drifted",
			combo.Different,
		},
		{
			"managed by $iac (equal IDs)",
			combo.Equal,
		},
		{
			"managed by $iac (equal IDs & attributes)",
			combo.DeepEqual,
		},
	} {
		l := len(data.list)
		if l == 0 {
			continue
		}
		ttl := strings.ReplaceAll(data.title, "$iac", data.list[0].IAC)
		lines = append(lines, fmt.Sprintf("Found resources %s", ttl))
		resTotal := 0
		for _, res := range data.list {
			resTotal += len(res.ResourceIDs)
			lines = append(lines, fmt.Sprintf("  %s:%s:", res.Provider, res.ResourceType))
			for _, id := range res.ResourceIDs {
				lines = append(lines, fmt.Sprintf("    - %s", id))
			}
		}
		total += resTotal
		summary = append(summary, fmt.Sprintf(" - %d %s", resTotal, ttl))
	}

	if len(lines) == 0 {
		return "No results"
	}

	lines = append(lines, fmt.Sprintf("Found %d resource(s)", total))

	covered := float64(len(combo.Equal) + len(combo.DeepEqual) + len(combo.Different)) // one of Equal and DeepEqual is supposed to be 0 depending on deep flag
	lines = append(lines, fmt.Sprintf(" - %.0f%% coverage", covered/float64(total)*100))
	lines = append(lines, summary...)

	return strings.Join(lines, "\n")
}
