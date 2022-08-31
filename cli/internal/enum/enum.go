package enum

import (
	"fmt"
	"strings"
)

// https://github.com/spf13/pflag/issues/236#issuecomment-931600452

type Enum struct {
	Allowed []string
	Value   string
}

// newEnum give a list of allowed flag parameters, where the second argument is the default
func NewEnum(allowed []string, d string) *Enum {
	return &Enum{
		Allowed: allowed,
		Value:   d,
	}
}

func (a Enum) String() string {
	return a.Value
}

func (a *Enum) Set(p string) error {
	isIncluded := func(opts []string, val string) bool {
		for _, opt := range opts {
			if val == opt {
				return true
			}
		}
		return false
	}
	if !isIncluded(a.Allowed, p) {
		return fmt.Errorf("%s is not included in %s", p, strings.Join(a.Allowed, ","))
	}
	a.Value = p
	return nil
}

func (*Enum) Type() string {
	return "string"
}
