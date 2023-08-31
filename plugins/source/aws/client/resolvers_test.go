package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/scalar"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	types1 "github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	types2 "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/assert"
)

func TestResolveTags(t *testing.T) {
	cases := []struct {
		InputItem    any
		ExpectedTags map[string]string
	}{
		{
			InputItem: types1.ListWebhookItem{ // non-ptr
				Tags: []types1.Tag{
					{
						Key:   aws.String("k1"),
						Value: aws.String("v1"),
					},
				},
			},
			ExpectedTags: map[string]string{"k1": "v1"},
		},
		{
			InputItem: &types2.EventSubscription{ // ptr
				Tags: []types2.Tag{
					{
						Key:   aws.String("k2"),
						Value: aws.String("v2"),
					},
				},
			},
			ExpectedTags: map[string]string{"k2": "v2"},
		},
		{
			InputItem: types1.ListWebhookItem{ // non-ptr, nil
				Tags: nil,
			},
			ExpectedTags: map[string]string{},
		},
	}

	for _, tc := range cases {
		ta := &schema.Table{
			Columns: []schema.Column{
				{
					Name: "tags",
					Type: sdkTypes.ExtensionTypes.JSON,
				},
			},
		}
		r := schema.NewResourceData(ta, nil, tc.InputItem)
		err := ResolveTags(context.Background(), nil, r, ta.Columns[0])
		assert.NoError(t, err)
		expectedJson := &scalar.JSON{}
		_ = expectedJson.Set(tc.ExpectedTags)
		assert.Equal(t, expectedJson, r.Get(ta.Columns[0].Name))
	}
}
