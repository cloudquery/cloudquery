package recipes

import (
	"github.com/google/go-github/v48/github"
)

func Repositories() []*Resource {
	repo := repository()
	repo.Service = "repositories"
	repo.TableName = "repositories"
	repo.Multiplex = orgMultiplex

	return []*Resource{repo}
}

func repository() *Resource {
	const (
		createdAt = "CreatedAt"
		pushedAt  = "PushedAt"
		updatedAt = "UpdatedAt"
	)

	return &Resource{
		SubService: "repositories",
		Struct:     new(github.Repository),
		SkipFields: append(skipID,
			createdAt, pushedAt, updatedAt,
		),
		ExtraColumns: append(orgColumns, idColumn,
			timestampField("created_at", createdAt),
			timestampField("pushed_at", pushedAt),
			timestampField("updated_at", updatedAt),
		),
	}
}
