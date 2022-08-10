package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTriggersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var name string
	require.NoError(t, faker.FakeData(&name))
	m.EXPECT().ListTriggers(
		gomock.Any(),
		&glue.ListTriggersInput{MaxResults: aws.Int32(200)},
	).Return(
		&glue.ListTriggersOutput{TriggerNames: []string{name}},
		nil,
	)

	var tr types.Trigger
	require.NoError(t, faker.FakeData(&tr))
	tr.Name = &name
	m.EXPECT().GetTrigger(
		gomock.Any(),
		&glue.GetTriggerInput{Name: aws.String(name)},
	).Return(
		&glue.GetTriggerOutput{Trigger: &tr},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: aws.String("arn:aws:glue:us-east-1:testAccount:trigger/" + name)},
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return client.Services{
		Glue: m,
	}
}

func TestTriggers(t *testing.T) {
	client.AwsMockTestHelper(t, Triggers(), buildTriggersMock, client.TestOptions{})
}
