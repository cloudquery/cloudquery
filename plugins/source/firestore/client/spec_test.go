package client

import (
	"encoding/base64"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpecValidate(t *testing.T) {
	testCases := []struct {
		name string
		spec Spec
		err  bool
	}{
		{
			name: "empty",
			spec: Spec{},
		},
		{
			name: "valid",
			spec: Spec{
				ProjectID:          "project_id",
				ServiceAccountJSON: "{}",
				OrderDirection:     "asc",
				MaxBatchSize:       1,
			},
		},
		{
			name: "invalid order direction",
			spec: Spec{
				ProjectID:          "project_id",
				ServiceAccountJSON: "{}",
				OrderDirection:     "invalid",
				MaxBatchSize:       1,
			},
			err: true,
		},
		{
			name: "invalid max batch size",
			spec: Spec{
				ProjectID:          "project_id",
				ServiceAccountJSON: "{}",
				OrderDirection:     "asc",
				MaxBatchSize:       -1,
			},
			err: true,
		},
		{
			name: "base64 service account",
			spec: Spec{
				ProjectID:          "project_id",
				ServiceAccountJSON: base64.StdEncoding.EncodeToString([]byte("{}")),
				OrderDirection:     "asc",
				MaxBatchSize:       1,
				UseBase64:          true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.spec.Validate()
			if tc.err && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tc.err && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
		},
		{
			Name: "valid spec",
			Spec: `{"project_id": "project_id","service_account_json": "{}","order_direction": "asc","max_batch_size": 1}`,
		},
		{
			Name: "invalid order direction",
			Spec: `{"project_id": "project_id","service_account_json": "{}","order_direction": "invalid","max_batch_size": 1}`,
			Err:  true,
		},
		{
			Name: "invalid max batch size",
			Spec: `{"project_id": "project_id","service_account_json": "{}","order_direction": "asc","max_batch_size": -1}`,
			Err:  true,
		},
		{
			Name: "base64 service account",
			Spec: `{"project_id": "project_id","service_account_json": "` + base64.StdEncoding.EncodeToString([]byte("{}")) + `","order_direction": "asc","max_batch_size": 1,"use_base64": true}`,
		},
	})
}
