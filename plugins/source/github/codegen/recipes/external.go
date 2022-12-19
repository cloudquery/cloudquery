package recipes

import (
	"github.com/google/go-github/v48/github"
)

func External() []*Resource {
	const (
		groupID   = "GroupID"
		updatedAt = "UpdatedAt"
	)
	return []*Resource{
		{
			Service:    "external",
			SubService: "groups",
			Multiplex:  orgMultiplex,
			Struct:     new(github.ExternalGroup),
			SkipFields: append(skipID, groupID, updatedAt),
			ExtraColumns: append(orgColumns,
				pkColumn("group_id", groupID),
				timestampField("updated_at", updatedAt),
			),
		},
	}
}
