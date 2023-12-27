package dms

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/thoas/go-funk"
)

func TestGetTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockDatabasemigrationserviceClient(ctrl)

	arn1, arn2 := "arn1", "arn2"

	m.EXPECT().ListTagsForResource(gomock.Any(), &databasemigrationservice.ListTagsForResourceInput{ResourceArnList: []string{arn1, arn2}}, gomock.Any()).Return(
		&databasemigrationservice.ListTagsForResourceOutput{
			TagList: []types.Tag{
				{ResourceArn: &arn1, Key: aws.String("tag1"), Value: aws.String("value1")},
				{ResourceArn: &arn2, Key: aws.String("tag2"), Value: aws.String("value2")},
				{ResourceArn: &arn2, Key: aws.String("tag3"), Value: aws.String("value3")},
			},
		}, nil)

	type model1 struct {
		Arn *string
	}
	models := []model1{
		{Arn: &arn1},
		{Arn: &arn2},
	}

	result, err := getTags(context.Background(), m, models, "Arn", func(options *databasemigrationservice.Options) {
		options.Region = "us-east-1"
	})
	require.NoError(t, err)

	require.Equal(t, map[string]map[string]any{
		arn1: {"tag1": "value1"},
		arn2: {"tag2": "value2", "tag3": "value3"},
	}, result)
}

func TestPutTags(t *testing.T) {
	type Model struct {
		Arn  *string
		Tags map[string]any
	}

	theArn := "test-arn"
	otherArn := "other-arn"
	anotherArn := "another-arn"
	data := []Model{
		{Arn: &theArn},
		{Arn: &otherArn},
		{Arn: &anotherArn},
	}

	tags := map[string]map[string]any{
		theArn:             {"test-key": "test-value"},
		"non-existent-arn": {"test-key2": "test-value2"},
		anotherArn:         {"test3-key": "test3-value"},
	}
	require.NoError(t, putTags(data, tags, "Arn"))

	expectTags := []any{
		tags[theArn],
		nil,
		tags[anotherArn],
	}

	for i := range expectTags {
		result := funk.Get(data[i], "Tags")
		require.Equalf(t, expectTags[i], result, "index %d", i)
	}
}

func TestPutTagsWrapped(t *testing.T) {
	type Model struct {
		Arn *string
	}

	type WrappedModel struct {
		Model
		Tags map[string]any
	}

	theArn := "test-arn"
	otherArn := "other-arn"
	anotherArn := "another-arn"
	data := []WrappedModel{
		{Model: Model{Arn: &theArn}},
		{Model: Model{Arn: &otherArn}},
		{Model: Model{Arn: &anotherArn}},
	}

	tags := map[string]map[string]any{
		theArn:             {"test-key": "test-value"},
		"non-existent-arn": {"test-key2": "test-value2"},
		anotherArn:         {"test3-key": "test3-value"},
	}
	require.NoError(t, putTags(data, tags, "Arn"))

	expectTags := []any{
		tags[theArn],
		nil,
		tags[anotherArn],
	}

	for i := range expectTags {
		result := funk.Get(data[i], "Tags")
		require.Equalf(t, expectTags[i], result, "index %d", i)
	}
}
