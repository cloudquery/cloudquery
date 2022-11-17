package resource

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

type Parent struct {
	Resource *Resource
	Path     string

	forceRequired bool
	required      []string
}

func (p *Parent) Parent() *Parent {
	switch {
	case p == nil:
		return nil
	case p.Resource.parent == nil:
		return nil
	default:
		return &Parent{
			Resource: p.Resource.parent,
			Path:     p.Path + ".Parent",
			required: p.required,
		}
	}
}

func (p *Parent) Required() bool {
	return p.forceRequired ||
		slices.Contains(p.required, strcase.ToLowerCamel(pluralize.NewClient().Singular(p.Resource.SubService)))
}

// Parent returns the parent resource, if any
func (r *Resource) Parent() *Parent {
	if r.parent == nil {
		return nil
	}

	return &Parent{
		Resource:      r.parent,
		forceRequired: true,
		required:      extractParamNames(r.fetcher.Params),
	}
}

func extractParamNames(params []string) (names []string) {
	if len(params) == 0 {
		return nil
	}
	names = make([]string, 0, len(params))
	for _, param := range params {
		names = append(names, extractParamName(param))
	}
	return
}

func extractParamName(param string) string {
	return strings.TrimPrefix(strings.Split(param, ".")[0], "*")
}
