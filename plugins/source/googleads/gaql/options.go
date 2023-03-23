package gaql

import (
	"golang.org/x/exp/slices"
)

type Options struct {
	Skip   []string
	Expand []string
}

func (o *Options) trim(prefix string) *Options {
	return &Options{
		Skip:   trimStrings(prefix, o.Skip),
		Expand: trimStrings(prefix, o.Expand),
	}
}

func (o *Options) skip(name string) bool {
	if o == nil {
		return false
	}
	return slices.Contains(o.Skip, name)
}

func (o *Options) expand(name string) bool {
	if o == nil {
		return false
	}
	return slices.Contains(o.Expand, name)
}
