package client

import (
	"crypto/sha256"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

const hashColumnName = "_cq_pk_hash_uuid"

func hashUUID(table *schema.Table) func(item []any) string {
	names := table.PrimaryKeys()
	if len(names) == 0 {
		return func([]any) string { return uuid.New().String() }
	}

	idx := make(map[string]int)
	for _, name := range names {
		idx[name] = table.Columns.Index(name)
	}

	return func(item []any) string {
		h := sha256.New()
		for name, i := range idx {
			h.Write([]byte(name))
			h.Write([]byte(fmt.Sprint(item[i])))
		}
		return uuid.NewSHA1(uuid.UUID{}, h.Sum(nil)).String()
	}
}
