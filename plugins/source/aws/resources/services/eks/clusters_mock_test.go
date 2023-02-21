package eks

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEksClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEksClient(ctrl)
	l := eks.DescribeClusterOutput{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListClustersOutput{
			Clusters: []string{"test-cluster"},
		}, nil)
	m.EXPECT().DescribeCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&l, nil)

	fp := types.FargateProfile{}
	err = faker.FakeObject(&fp)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListFargateProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListFargateProfilesOutput{
			FargateProfileNames: []string{"test-profile"},
		}, nil)
	m.EXPECT().DescribeFargateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeFargateProfileOutput{FargateProfile: &fp}, nil)

	ng := types.Nodegroup{}
	err = faker.FakeObject(&ng)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListNodegroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListNodegroupsOutput{
			Nodegroups: []string{"test-nodegroup"},
		}, nil)
	m.EXPECT().DescribeNodegroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeNodegroupOutput{
			Nodegroup: &ng,
		}, nil)
	return client.Services{
		Eks: m,
	}
}

func TestEksClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEksClusters, client.TestOptions{})
}
