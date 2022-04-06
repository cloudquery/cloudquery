package drift

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/stretchr/testify/assert"
)

func TestParseTerraformInstance(t *testing.T) {
	cases := []struct {
		TFAttrs     map[string]interface{}
		Identifiers []string
		CloudAttrs  AttrList
		Path        string
		ExpectedIDs []string
	}{
		{
			TFAttrs: map[string]interface{}{
				"id":    "the-id",
				"value": "the-value",
			},
			Identifiers: []string{"id"},
			ExpectedIDs: []string{"the-id"},
		},
		{
			TFAttrs: map[string]interface{}{
				"arn": "the-arn",
				"params": []map[string]interface{}{
					{"val": "a", "other-key": "xx"},
					{"val": "b", "other-key": "yy"},
				},
				"value": "the-value",
			},
			Identifiers: []string{"arn", "params.0.val"},
			ExpectedIDs: []string{"the-arn|a"},
		},
		{
			TFAttrs: map[string]interface{}{
				"arn": "the-arn",
				"params": []map[string]interface{}{
					{"val": "a", "other-key": "xx"},
					{"val": "b", "other-key": "yy"},
				},
				"value": "the-value",
			},
			Identifiers: []string{"arn", "params.#.val"},
			ExpectedIDs: []string{"the-arn|a", "the-arn|b"},
		},
		{
			TFAttrs: map[string]interface{}{
				"arn": "the-arn",
				"params": []map[string]interface{}{
					{"val": "a", "other-key": "xx"},
					{"val": "b", "other-key": "yy"},
				},
				"value": "the-value",
			},
			Path:        "params",
			Identifiers: []string{"root.arn", "val"},
			ExpectedIDs: []string{"the-arn|a", "the-arn|b"},
		},
		{
			TFAttrs: map[string]interface{}{
				"arn": "the-arn",
				"data": []map[string]interface{}{
					{"sub-key": "the-sub-val1"},
					{"sub-key": "the-sub-val2"},
				},
				"value": "the-value",
			},
			Path:        "data",
			Identifiers: []string{"root.arn", "sub-key"},
			ExpectedIDs: []string{"the-arn|the-sub-val1", "the-arn|the-sub-val2"},
		},
	}
	for _, tc := range cases {
		raw, _ := json.Marshal(tc.TFAttrs)

		resList := parseTerraformInstance(
			terraform.Instance{
				AttributesRaw: raw,
			},
			tc.Identifiers,
			tc.CloudAttrs,
			tc.Path,
		)
		assert.EqualValues(t, tc.ExpectedIDs, resList.IDs())
	}
}
