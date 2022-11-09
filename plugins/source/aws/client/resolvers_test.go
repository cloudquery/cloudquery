package client

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	types1 "github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	types2 "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/assert"
)

type SliceJsonStruct struct {
	Value  []types1.Tag
	Nested *SliceJsonStruct
}

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
		r := schema.NewResourceData(ta, nil, tc.InputItem)
		err := ResolveTags(context.Background(), nil, r, ta.Columns[0])
		assert.NoError(t, err)
		expectedJson := &schema.JSON{}
		_ = expectedJson.Set(tc.ExpectedTags)
		assert.Equal(t, expectedJson, r.Get(ta.Columns[0].Name))
	}
}

func TestResolveSliceJson(t *testing.T) {
	cases := []struct {
		InputItem    interface{}
		ExpectedData map[string]interface{}
		path         string
		keyPath      string
		valuePath    string
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
			ExpectedData: map[string]interface{}{"k1": aws.String("v1")},
			path:         "Tags",
			keyPath:      "Key",
			valuePath:    "Value",
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
			ExpectedData: map[string]interface{}{"k2": aws.String("v2")},
			path:         "Tags",
			keyPath:      "Key",
			valuePath:    "Value",
		},
		{
			InputItem: SliceJsonStruct{Nested: &SliceJsonStruct{
				Nested: &SliceJsonStruct{
					Value: []types1.Tag{{
						Key:   aws.String("k1"),
						Value: aws.String("v1"),
					}, {
						Key:   aws.String("k2"),
						Value: aws.String("v2"),
					}},
				},
			}},
			ExpectedData: map[string]interface{}{"k1": aws.String("v1"), "k2": aws.String("v2")},
			path:         "Nested.Nested.Value",
			keyPath:      "Key",
			valuePath:    "Value",
		},
		{
			InputItem: types1.ListWebhookItem{ // non-ptr, nil
				Tags: nil,
			},
			ExpectedData: nil,
			path:         "Tags",
			keyPath:      "Key",
			valuePath:    "Value",
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
		r := schema.NewResourceData(ta, nil, tc.InputItem)
		err := SliceJsonResolver(tc.path, tc.keyPath, tc.valuePath)(context.Background(), nil, r, ta.Columns[0])
		assert.NoError(t, err)
		expectedJson := &schema.JSON{}
		_ = expectedJson.Set(tc.ExpectedData)
		assert.Equal(t, expectedJson, r.Get(ta.Columns[0].Name))
	}
}

var jsonString = "{\"k1\":\"v1\"}"
var jsonBytes = []byte(jsonString)

func TestResolveStringJson(t *testing.T) {
	cases := []struct {
		InputItem    interface{}
		ExpectedData interface{}
		Path         string
	}{
		{
			InputItem: struct {
				Json string
			}{Json: jsonString},
			ExpectedData: map[string]interface{}{"k1": "v1"},
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json string
			}{Json: ""},
			ExpectedData: nil,
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json *string
			}{Json: &jsonString},
			ExpectedData: map[string]interface{}{"k1": "v1"},
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json *string
			}{Json: nil},
			ExpectedData: nil,
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json []byte
			}{Json: jsonBytes},
			ExpectedData: map[string]interface{}{"k1": "v1"},
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json *[]byte
			}{Json: &jsonBytes},
			ExpectedData: map[string]interface{}{"k1": "v1"},
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json *[]byte
			}{Json: nil},
			ExpectedData: nil,
			Path:         "Json",
		},
		{
			InputItem: struct {
				Json []byte
			}{Json: []byte{}},
			ExpectedData: nil,
			Path:         "Json",
		},
	}

	for _, tc := range cases {
		ta := &schema.Table{
			Columns: []schema.Column{
				{
					Name: "json",
					Type: schema.TypeJSON,
				},
			},
		}
		r := schema.NewResourceData(ta, nil, tc.InputItem)
		err := MarshaledJsonResolver(tc.Path)(context.Background(), nil, r, ta.Columns[0])
		assert.NoError(t, err)
		expectedJson := &schema.JSON{}
		_ = expectedJson.Set(tc.ExpectedData)
		assert.Equal(t, expectedJson, r.Get(ta.Columns[0].Name))
	}
}
