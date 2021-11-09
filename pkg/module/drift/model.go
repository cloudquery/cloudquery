package drift

import (
	"fmt"
	"sort"
	"strings"
)

type RunParams struct {
	Debug bool

	TfMode      string
	ForceDeep   bool
	ListManaged bool
	AccountIDs  []string

	IACName    string
	StateFiles []string
}

type Resource struct {
	ID         string        `json:"id"`
	Attributes []interface{} `json:"-"`
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

func (r ResourceList) Walk(fn func(*Resource)) {
	for i := range r {
		fn(r[i])
	}
}

// Map returns a map of ID vs. attributes
func (r ResourceList) Map() map[string][]interface{} {
	ret := make(map[string][]interface{}, len(r))
	for i := range r {
		ret[r[i].ID] = r[i].Attributes
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

type Results struct {
	Data []*Result `json:"data"`

	// Options
	ListManaged bool `json:"-"` // Show or hide Equal/DeepEqual output
	Debug       bool `json:"-"` // Print debug output regarding results

	// These fields are calculated
	Drifted  int     `json:"drifted_res"`
	Covered  int     `json:"covered_res"`
	Total    int     `json:"total_res"`
	Coverage float64 `json:"coverage_pct"`

	Text string `json:"-"`
}

func (rs *Results) String() string {
	return rs.Text
}

func (rs *Results) ExitCode() int {
	if rs.Drifted > 0 {
		return 1
	}
	return 0
}

func (rs *Results) process() {
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

	for _, r := range rs.Data {
		if r == nil {
			continue
		}
		transform(r, r.Different, &combo.Different)
		transform(r, r.Extra, &combo.Extra)
		transform(r, r.Equal, &combo.Equal)
		transform(r, r.DeepEqual, &combo.DeepEqual)
		transform(r, r.Missing, &combo.Missing)
	}

	var ( // nolint: prealloc
		lines   []string
		summary []string
	)

	for _, data := range []struct {
		title       string
		list        []combined
		hideListing bool
		drift       bool
	}{
		{
			"not managed by $iac",
			combo.Extra,
			false,
			true,
		},
		{
			"in $iac state but missing on the cloud provider",
			combo.Missing,
			false,
			true,
		},
		{
			"managed by $iac but drifted",
			combo.Different,
			false,
			true,
		},
		{
			"managed by $iac (equal IDs)",
			combo.Equal,
			!rs.ListManaged,
			false,
		},
		{
			"managed by $iac (equal IDs & attributes)",
			combo.DeepEqual,
			!rs.ListManaged,
			false,
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
			if data.hideListing {
				continue
			}

			lines = append(lines, fmt.Sprintf("  %s:%s:", res.Provider, res.ResourceType))
			for _, id := range res.ResourceIDs {
				lines = append(lines, fmt.Sprintf("    - %s", id))
			}
		}
		if data.hideListing {
			lines[len(lines)-1] += fmt.Sprintf(" (%d)", resTotal) // append count to previous line
		}
		if data.drift {
			rs.Drifted += resTotal
		}

		rs.Total += resTotal
		summary = append(summary, fmt.Sprintf(" - %d %s", resTotal, ttl))
	}

	if len(lines) == 0 {
		rs.Text = "No results"
		return
	}

	lines = append(lines, fmt.Sprintf("Found %d resource(s)", rs.Total))

	// one of Equal and DeepEqual is supposed to be 0 depending on deep flag
	for _, l := range [][]combined{combo.Equal, combo.DeepEqual, combo.Different} {
		for _, z := range l {
			rs.Covered += len(z.ResourceIDs)
		}
	}

	rs.Coverage = float64(rs.Covered) / float64(rs.Total)
	cvg := fmt.Sprintf("%.2f", rs.Coverage*100)
	cvg = strings.ReplaceAll(cvg, ".00", "")

	lines = append(lines, fmt.Sprintf(" - %s%% coverage", cvg))
	lines = append(lines, summary...)

	if rs.Debug {
		matchedResourceTypes := make(map[string]struct{})
		for _, t := range append(append(combo.Equal, combo.DeepEqual...), combo.Different...) {
			matchedResourceTypes[t.Provider+":"+t.ResourceType] = struct{}{}
		}
		var unmatchedTypes []string
		for _, t := range append(combo.Extra, combo.Missing...) {
			k := t.Provider + ":" + t.ResourceType
			if _, ok := matchedResourceTypes[k]; ok {
				continue
			}
			unmatchedTypes = append(unmatchedTypes, k)
		}
		if len(unmatchedTypes) > 0 {
			sort.Strings(unmatchedTypes)
			lines = append(lines, "These types weren't matched: "+strings.Join(unmatchedTypes, ", "))
		}
	}

	rs.Text = strings.Join(lines, "\n")
}
