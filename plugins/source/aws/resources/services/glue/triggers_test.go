package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTriggersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var name string
	require.NoError(t, faker.FakeObject(&name))
	m.EXPECT().ListTriggers(
		gomock.Any(),
		&glue.ListTriggersInput{MaxResults: aws.Int32(200)},
		gomock.Any(),
	).Return(
		&glue.ListTriggersOutput{TriggerNames: []string{name}},
		nil,
	)

	var tr types.Trigger
	require.NoError(t, faker.FakeObject(&tr))
	tr.Name = &name
	m.EXPECT().GetTrigger(
		gomock.Any(),
		&glue.GetTriggerInput{Name: aws.String(name)},
		gomock.Any(),
	).Return(
		&glue.GetTriggerOutput{Trigger: &tr},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: aws.String("arn:aws:glue:us-east-1:testAccount:trigger/" + name)},
		gomock.Any(),
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
