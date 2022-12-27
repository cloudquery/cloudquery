package recipes

import (
	"github.com/google/go-github/v48/github"
)

func Dependabot() []*Resource {
	return []*Resource{}
}

func alerts() *Resource {
	return &Resource{
		SubService:           "alerts",
		Struct:               new(github.DependabotAlert),
		SkipFields:           nil,
		PKColumns:            nil,
		ExtraColumns:         nil,
		Table:                nil,
		TableName:            "",
		Multiplex:            "",
		PreResourceResolver:  "",
		PostResourceResolver: "",
		Relations:            nil,
	}
}
