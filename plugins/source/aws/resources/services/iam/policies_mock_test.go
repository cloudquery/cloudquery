package iam

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIamPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Policy{}
	require.NoError(t, faker.FakeObject(&g))
	g.Arn = aws.String("arn:aws:iam::testAccount:policy/IAMReadOnlyAccess")
	m.EXPECT().ListPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListPoliciesOutput{
			Policies: []iamTypes.Policy{g},
		}, nil)

	tag := iamTypes.Tag{}
	require.NoError(t, faker.FakeObject(&tag))
	m.EXPECT().ListPolicyTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListPolicyTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)

	createDate := time.Now()
	m.EXPECT().ListPolicyVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListPolicyVersionsOutput{
			Versions: []iamTypes.PolicyVersion{{
				CreateDate: &createDate,
				VersionId:  aws.String("v1"),
			}},
		},
		nil,
	)

	m.EXPECT().GetPolicyVersion(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetPolicyVersionOutput{
			PolicyVersion: &iamTypes.PolicyVersion{
				VersionId:  aws.String("v1"),
				Document:   aws.String(`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"iam:*","Resource":"*"}]}`),
				CreateDate: &createDate,
			},
		},
		nil,
	)

	m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.GenerateServiceLastAccessedDetailsOutput{JobId: aws.String("JobId")}, nil)

	lastAccessed := []iamTypes.ServiceLastAccessed{}
	require.NoError(t, faker.FakeObject(&lastAccessed))
	m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetServiceLastAccessedDetailsOutput{ServicesLastAccessed: lastAccessed, JobStatus: iamTypes.JobStatusTypeCompleted},
		nil,
	)

	return client.Services{
		Iam: m,
	}
}

func TestIamPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, Policies(), buildIamPolicies, client.TestOptions{})
}
