package drift

import (
	"fmt"
	"strings"
)

type Resource struct {
	ID string
}

type Result struct {
	Provider     string
	ResourceType string

	Different []*Resource
	Equal     []*Resource
	Missing   []*Resource
	Extra     []*Resource
}

func (r *Result) String() string {
	var parts []string
	if l := len(r.Different); l > 0 {
		parts = append(parts, fmt.Sprintf("%d different", l))
	}
	if l := len(r.Equal); l > 0 {
		parts = append(parts, fmt.Sprintf("%d equal", l))
	}
	if l := len(r.Missing); l > 0 {
		parts = append(parts, fmt.Sprintf("%d missing", l))
	}
	if l := len(r.Extra); l > 0 {
		parts = append(parts, fmt.Sprintf("%d extra", l))
	}
	if len(parts) == 0 {
		parts = append(parts, "no")
	}

	return fmt.Sprintf("for %s:%s we have %s resources", r.Provider, r.ResourceType, strings.Join(parts, ", "))
}

func (r *Result) HasResources() bool {
	return len(r.Different)+len(r.Equal)+len(r.Missing)+len(r.Extra) > 0
}

type Results []*Result

func (rs Results) String() string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		if !r.HasResources() {
			continue
		}
		parts = append(parts, r.String())
	}

	if len(parts) == 0 {
		return "no results"
	}

	return strings.Join(parts, "\n")
}
