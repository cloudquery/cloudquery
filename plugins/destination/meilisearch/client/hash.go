package client

import (
	"crypto/sha256"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

const hashColumnName = "_cq_pk_hash_uuid"

func hashUUID(sc *arrow.Schema) func(map[string]any) string {
	pk := schema.PrimaryKeyIndices(sc)
	if len(pk) == 0 {
		return func(map[string]any) string { return uuid.New().String() }
	}

	names := make([]string, len(pk))
	for i, idx := range pk {
		names[i] = sc.Field(idx).Name
	}

	return func(row map[string]any) string {
		h := sha256.New()
		for _, name := range names {
			h.Write([]byte(name))
			h.Write([]byte(fmt.Sprint(row[name])))
		}
		return uuid.NewSHA1(uuid.UUID{}, h.Sum(nil)).String()
	}
}
