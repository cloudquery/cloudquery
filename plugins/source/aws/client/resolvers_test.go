package client

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	types1 "github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	types2 "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/stretchr/testify/assert"
)

func TestResolveTags(t *testing.T) {
	cases := []struct {
		InputItem    interface{}
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
					Type: schema.TypeJSON,
				},
			},
		}
		r := schema.NewResourceData(schema.PostgresDialect{}, ta, nil, tc.InputItem, nil, time.Now())
		err := ResolveTags(context.Background(), nil, r, ta.Columns[0])
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedTags, r.Get(ta.Columns[0].Name))
	}
}
