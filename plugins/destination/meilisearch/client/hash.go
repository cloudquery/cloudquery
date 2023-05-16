package client

import (
	"crypto/sha256"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
)

const hashColumnName = "_cq_pk_hash_uuid"

func hashUUID(table *schema.Table) func(map[string]any) string {
	pk := table.PrimaryKeys()
	if len(pk) == 0 {
		return func(map[string]any) string { return uuid.New().String() }
	}

	return func(row map[string]any) string {
		h := sha256.New()
		for _, name := range pk {
			h.Write([]byte(name))
			h.Write([]byte(fmt.Sprint(row[name])))
		}
		return uuid.NewSHA1(uuid.UUID{}, h.Sum(nil)).String()
	}
}
