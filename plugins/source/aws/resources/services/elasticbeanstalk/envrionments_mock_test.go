package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	require.NoError(t, faker.FakeObject(&la))

	l := elasticbeanstalkTypes.EnvironmentDescription{
		ApplicationName: la.ApplicationName,
	}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)

	tags := elasticbeanstalk.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	configSettingsOutput := elasticbeanstalk.DescribeConfigurationSettingsOutput{}
	require.NoError(t, faker.FakeObject(&configSettingsOutput))
	m.EXPECT().DescribeConfigurationSettings(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configSettingsOutput, nil)

	configOptsOutput := elasticbeanstalk.DescribeConfigurationOptionsOutput{}
	require.NoError(t, faker.FakeObject(&configOptsOutput))
	m.EXPECT().DescribeConfigurationOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configOptsOutput, nil)

	return client.Services{
		Elasticbeanstalk: m,
	}
}

func TestElasticbeanstalkEnvironments(t *testing.T) {
	client.AwsMockTestHelper(t, Environments(), buildElasticbeanstalkEnvironments, client.TestOptions{})
}
