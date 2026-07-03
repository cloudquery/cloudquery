package client

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Mirrors k8s_core_pods.status_container_statuses: a JSON column (list of structs)
// whose TypeSchema marks nested finishedAt as a timestamp. The transform must emit
// a real time.Time for it, not the RFC3339 string, while leaving other leaves untouched.
func TestTransformRecord_ConvertsNestedTimestampsUsingTypeSchema(t *testing.T) {
	jb := types.NewJSONBuilder(memory.DefaultAllocator)
	jb.Append([]any{
		map[string]any{
			"lastState": map[string]any{
				"terminated": map[string]any{
					"finishedAt": "2026-06-13T08:36:36Z",
					"reason":     "Completed",
				},
			},
		},
	})
	arr := jb.NewJSONArray()
	defer arr.Release()

	sc := arrow.NewSchema([]arrow.Field{{Name: "status_container_statuses", Type: types.ExtensionTypes.JSON}}, nil)
	rec := array.NewRecordBatch(sc, []arrow.Array{arr}, 1)
	defer rec.Release()

	table := &schema.Table{
		Name: "k8s_core_pods",
		Columns: schema.ColumnList{
			{
				Name:       "status_container_statuses",
				Type:       types.ExtensionTypes.JSON,
				TypeSchema: `[{"lastState":{"terminated":{"finishedAt":"timestamp[us, tz=UTC]","reason":"utf8"}}}]`,
			},
		},
	}

	docs := (&Client{}).transformRecord(table, rec)
	require.Len(t, docs, 1)

	statuses := docs[0].(bson.M)["status_container_statuses"].([]any)
	require.Len(t, statuses, 1)
	terminated := statuses[0].(map[string]any)["lastState"].(map[string]any)["terminated"].(map[string]any)

	finishedAt, ok := terminated["finishedAt"].(time.Time)
	require.Truef(t, ok, "finishedAt should be time.Time, got %T (%v)", terminated["finishedAt"], terminated["finishedAt"])
	require.Equal(t, "2026-06-13T08:36:36Z", finishedAt.UTC().Format(time.RFC3339))

	// a non-timestamp leaf must remain a plain string
	require.Equal(t, "Completed", terminated["reason"])
}

// Determinism guard: only leaves the schema types as timestamps convert. A string that
// happens to look like a date but is typed utf8 must stay a string (no heuristic guessing).
func TestTransformRecord_OnlyConvertsTimestampTypedLeaves(t *testing.T) {
	jb := types.NewJSONBuilder(memory.DefaultAllocator)
	jb.Append(map[string]any{
		"finishedAt": "2026-06-13T08:36:36Z", // typed timestamp -> converts
		"version":    "2026-01-02T00:00:00Z", // date-looking but typed utf8 -> stays string
	})
	arr := jb.NewJSONArray()
	defer arr.Release()

	sc := arrow.NewSchema([]arrow.Field{{Name: "data", Type: types.ExtensionTypes.JSON}}, nil)
	rec := array.NewRecordBatch(sc, []arrow.Array{arr}, 1)
	defer rec.Release()

	table := &schema.Table{Columns: schema.ColumnList{{
		Name:       "data",
		Type:       types.ExtensionTypes.JSON,
		TypeSchema: `{"finishedAt":"timestamp[us, tz=UTC]","version":"utf8"}`,
	}}}

	data := (&Client{}).transformRecord(table, rec)[0].(bson.M)["data"].(map[string]any)
	_, isTime := data["finishedAt"].(time.Time)
	require.True(t, isTime, "timestamp-typed leaf should convert")
	require.Equal(t, "2026-01-02T00:00:00Z", data["version"], "utf8-typed date-looking string must stay a string")
}

// Without a TypeSchema the JSON column is passed through unchanged (best-effort gating).
func TestTransformRecord_NoTypeSchemaLeavesValuesUnchanged(t *testing.T) {
	jb := types.NewJSONBuilder(memory.DefaultAllocator)
	jb.Append(map[string]any{"finishedAt": "2026-06-13T08:36:36Z"})
	arr := jb.NewJSONArray()
	defer arr.Release()

	sc := arrow.NewSchema([]arrow.Field{{Name: "data", Type: types.ExtensionTypes.JSON}}, nil)
	rec := array.NewRecordBatch(sc, []arrow.Array{arr}, 1)
	defer rec.Release()

	table := &schema.Table{Columns: schema.ColumnList{{Name: "data", Type: types.ExtensionTypes.JSON}}}
	data := (&Client{}).transformRecord(table, rec)[0].(bson.M)["data"].(map[string]any)
	require.Equal(t, "2026-06-13T08:36:36Z", data["finishedAt"])
}

// A running container has terminated:null; walking the schema past a nil must not panic.
func TestTransformRecord_HandlesNullNestedValues(t *testing.T) {
	jb := types.NewJSONBuilder(memory.DefaultAllocator)
	jb.Append([]any{map[string]any{"lastState": map[string]any{"terminated": nil}}})
	arr := jb.NewJSONArray()
	defer arr.Release()

	sc := arrow.NewSchema([]arrow.Field{{Name: "status_container_statuses", Type: types.ExtensionTypes.JSON}}, nil)
	rec := array.NewRecordBatch(sc, []arrow.Array{arr}, 1)
	defer rec.Release()

	table := &schema.Table{Columns: schema.ColumnList{{
		Name:       "status_container_statuses",
		Type:       types.ExtensionTypes.JSON,
		TypeSchema: `[{"lastState":{"terminated":{"finishedAt":"timestamp[us, tz=UTC]"}}}]`,
	}}}

	statuses := (&Client{}).transformRecord(table, rec)[0].(bson.M)["status_container_statuses"].([]any)
	require.Nil(t, statuses[0].(map[string]any)["lastState"].(map[string]any)["terminated"])
}
