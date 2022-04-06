package drift

import (
	"fmt"
	"sort"
	"strings"

	cqsort "github.com/cloudquery/cloudquery/internal/sort"
)

type RunParams struct {
	Debug bool

	TfMode      string
	ForceDeep   bool
	ListManaged bool

	IACName    string
	StateFiles []string
}

type Resource struct {
	ID         string            `json:"id"`
	Attributes []interface{}     `json:"-"`
	Tags       map[string]string `json:"-"`
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

func (r ResourceList) Walk(fn func(*Resource), skipper func(*Resource) bool) {
	for i := range r {
		if skipper != nil && skipper(r[i]) {
			continue
		}
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
	IACName string    `json:"iac"`
	Data    []*Result `json:"data"`

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

	// transform appends the given resource id list of type provName and resType into the given slice of combined type, an item per provider and resource type.
	transform := func(ids []string, provName, resType string, dst *[]combined) {
		if len(ids) == 0 {
			return
		}

		// check if we already have the given resType in []combined somewhere, if so, append to that
		for i, c := range *dst {
			if c.Provider == provName && c.ResourceType == resType {
				(*dst)[i].ResourceIDs = append((*dst)[i].ResourceIDs, ids...)
				return
			}
		}

		*dst = append(*dst, combined{
			ResourceType: resType,
			Provider:     provName,
			ResourceIDs:  ids,
		})
	}

	for _, r := range rs.Data {
		if r == nil {
			continue
		}
		cleanRes, _ := SplitHashedResource(r.ResourceType)
		transform(r.Different.IDs(), r.Provider, cleanRes, &combo.Different)
		transform(r.Extra.IDs(), r.Provider, cleanRes, &combo.Extra)
		transform(r.Equal.IDs(), r.Provider, cleanRes, &combo.Equal)
		transform(r.DeepEqual.IDs(), r.Provider, cleanRes, &combo.DeepEqual)
		transform(r.Missing.IDs(), r.Provider, cleanRes, &combo.Missing)
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
		ttl := strings.ReplaceAll(data.title, "$iac", rs.IACName)

		resLines := make([]string, 0, l)
		resTotal := 0
		for _, res := range data.list {
			ids := cqsort.Unique(res.ResourceIDs)
			resTotal += len(ids)
			if data.hideListing {
				continue
			}

			resLines = append(resLines, fmt.Sprintf("  %s:%s:", res.Provider, res.ResourceType))
			for _, id := range ids {
				resLines = append(resLines, fmt.Sprintf("    - %s", id))
			}
		}

		lines = append(lines, fmt.Sprintf("%d Resources %s", resTotal, ttl))
		lines = append(lines, resLines...)

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

	summary = append([]string{fmt.Sprintf("Total number of resources: %d", rs.Total)}, summary...)

	// one of Equal and DeepEqual is supposed to be 0 depending on deep flag
	for _, l := range [][]combined{combo.Equal, combo.DeepEqual, combo.Different} {
		for _, z := range l {
			rs.Covered += len(cqsort.Unique(z.ResourceIDs))
		}
	}

	rs.Coverage = float64(rs.Covered) / float64(rs.Total)
	cvg := fmt.Sprintf("%.2f", rs.Coverage*100)
	cvg = strings.ReplaceAll(cvg, ".00", "")

	summary = append(summary, fmt.Sprintf(" - %s%% covered by %s", cvg, rs.IACName))

	lines = append([]string{"=== DRIFT RESULTS  ==="}, lines...)
	lines = append(lines, "=== SUMMARY ===")
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
