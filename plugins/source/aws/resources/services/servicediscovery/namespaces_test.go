package servicediscovery

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNamespaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockServicediscoveryClient(ctrl)

	var ns types.NamespaceSummary
	require.NoError(t, faker.FakeObject(&ns))
	m.EXPECT().ListNamespaces(
		gomock.Any(),
		&servicediscovery.ListNamespacesInput{MaxResults: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&servicediscovery.ListNamespacesOutput{Namespaces: []types.NamespaceSummary{ns}},
		nil,
	)

	var namespace types.Namespace
	require.NoError(t, faker.FakeObject(&namespace))
	namespace.Arn = ns.Arn
	namespace.Id = ns.Id

	m.EXPECT().GetNamespace(
		gomock.Any(),
		&servicediscovery.GetNamespaceInput{Id: ns.Id},
		gomock.Any(),
	).Return(
		&servicediscovery.GetNamespaceOutput{Namespace: &namespace},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&servicediscovery.ListTagsForResourceInput{ResourceARN: namespace.Arn},
		gomock.Any(),
	).Return(
		&servicediscovery.ListTagsForResourceOutput{Tags: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}}},
		nil,
	)

	return client.Services{
		Servicediscovery: m,
	}
}

func TestNamespaces(t *testing.T) {
	client.AwsMockTestHelper(t, Namespaces(), buildNamespaces, client.TestOptions{})
}
