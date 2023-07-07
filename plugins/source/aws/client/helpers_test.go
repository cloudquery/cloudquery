package client

import (
	"context"
	"errors"
	"testing"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	ttypes "github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveARN(t *testing.T) {
	// t *Table, parent *Resource, fetchTime time.Time, item interface{}
	tests := []struct {
		name       string
		columnName string
		service    AWSService
		resourceID func(resource *schema.Resource) ([]string, error)
		resource   *schema.Resource
		want       any
		wantErr    bool
	}{
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return []string{"restapis", *resource.Item.(types.RestApi).Id}, nil
			},
			schema.NewResourceData(&schema.Table{Columns: []schema.Column{{Name: "myarn", Type: arrow.BinaryTypes.String}}}, nil, types.RestApi{Id: aws.String("myid")}),
			&scalar.String{Valid: true, Value: "arn:aws:apigateway:region::restapis/myid"},
			false,
		},
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return []string{"", "restapis", *resource.Item.(types.RestApi).Id}, nil
			},
			schema.NewResourceData(&schema.Table{Columns: []schema.Column{{Name: "myarn", Type: arrow.BinaryTypes.String}}}, nil, types.RestApi{Id: aws.String("myid")}),
			&scalar.String{Valid: true, Value: "arn:aws:apigateway:region::/restapis/myid"},
			false,
		},
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return nil, errors.New("test")
			},
			schema.NewResourceData(&schema.Table{Columns: []schema.Column{{Name: "myarn", Type: arrow.BinaryTypes.String}}}, nil, types.RestApi{Id: aws.String("myid")}),
			&scalar.String{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := ResolveARN(tt.service, tt.resourceID)
			col := schema.Column{Name: tt.columnName}
			client := Client{Region: "region", Partition: "aws"}
			err := resolver(context.Background(), &client, tt.resource, col)

			actual := tt.resource.Get(tt.columnName)
			require.Equal(t, tt.want, actual)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestTagsToMap(t *testing.T) {
	type randomType struct {
		Key   string
		Value string
	}
	type randomTypePtr struct {
		Key   *string
		Value *string
	}

	tests := []struct {
		Input    any
		Expected map[string]string
	}{
		{
			Input: []randomType{
				{
					Key:   "k",
					Value: "v",
				},
			},
			Expected: map[string]string{"k": "v"},
		},
		{
			Input: []randomTypePtr{
				{
					Key:   aws.String("k"),
					Value: aws.String("v"),
				},
				{
					Key:   nil,
					Value: aws.String("emptykey"),
				},
				{
					Key:   aws.String("emptyvalue"),
					Value: nil,
				},
			},
			Expected: map[string]string{"k": "v", "emptyvalue": ""},
		},
		{
			Input: []ttypes.Tag{
				{
					Key:   nil,
					Value: nil,
				},
			},
			Expected: map[string]string{},
		},
		{
			Input: []ttypes.Tag{
				{
					Key:   aws.String(""),
					Value: aws.String(""),
				},
			},
			Expected: map[string]string{"": ""},
		},
		{
			Input: []ttypes.Tag{
				{
					Key:   aws.String("k"),
					Value: aws.String("v"),
				},
				{
					Key:   nil,
					Value: nil,
				},
				{
					Key:   aws.String("k2"),
					Value: aws.String("v2"),
				},
			},
			Expected: map[string]string{"k": "v", "k2": "v2"},
		},
	}
	for _, tc := range tests {
		res := TagsToMap(tc.Input)
		assert.Equal(t, tc.Expected, res)
	}
}

func TestTagsIntoMap(t *testing.T) {
	type randomType struct {
		Key   *string
		Value *string
	}

	res := TagsToMap([]randomType{
		{
			Key:   aws.String("k"),
			Value: aws.String("v"),
		},
	})

	assert.Equal(t, map[string]string{"k": "v"}, res)

	TagsIntoMap([]randomType{
		{
			Key:   aws.String("k2"),
			Value: aws.String("v2"),
		},
	}, res)

	assert.Equal(t, map[string]string{"k": "v", "k2": "v2"}, res)
}
