package client

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/stretchr/testify/require"
)

func TestResolveARN(t *testing.T) {
	tests := []struct {
		name       string
		columnName string
		service    AWSService
		resourceID func(resource *schema.Resource) ([]string, error)
		resource   *schema.Resource
		want       interface{}
		wantErr    bool
	}{
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return []string{"restapis", *resource.Item.(types.RestApi).Id}, nil
			},
			schema.NewResourceData(&schema.PostgresDialect{}, &schema.Table{Columns: []schema.Column{{Name: "myarn"}}}, nil, types.RestApi{Id: aws.String("myid")}, nil, time.Now()),
			"arn:aws:apigateway:region::restapis/myid",
			false,
		},
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return []string{"", "restapis", *resource.Item.(types.RestApi).Id}, nil
			},
			schema.NewResourceData(&schema.PostgresDialect{}, &schema.Table{Columns: []schema.Column{{Name: "myarn"}}}, nil, types.RestApi{Id: aws.String("myid")}, nil, time.Now()),
			"arn:aws:apigateway:region::/restapis/myid",
			false,
		},
		{
			"apigateway",
			"myarn",
			ApigatewayService,
			func(resource *schema.Resource) ([]string, error) {
				return nil, errors.New("test")
			},
			schema.NewResourceData(&schema.PostgresDialect{}, &schema.Table{Columns: []schema.Column{{Name: "myarn"}}}, nil, types.RestApi{Id: aws.String("myid")}, nil, time.Now()),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := ResolveARN(tt.service, tt.resourceID)
			col := schema.Column{Name: tt.columnName}
			client := Client{Region: "region"}
			err := resolver(context.Background(), &client, tt.resource, col)
			require.Equal(t, tt.resource.Get(tt.columnName), tt.want)
			require.Equal(t, err != nil, tt.wantErr)
		})
	}
}
