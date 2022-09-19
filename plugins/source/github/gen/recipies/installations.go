package recipies

import (
	"github.com/google/go-github/v45/github"
)

func Installations() []*Resource {
	return []*Resource{
		{
			Service:      "installations",
			SubService:   "installations",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Installation),
			TableName:    "installations",
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
		},
	}
}
