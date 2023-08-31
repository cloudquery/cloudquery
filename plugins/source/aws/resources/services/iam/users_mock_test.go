package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := types.User{}
	require.NoError(t, faker.FakeObject(&u))
	g := types.Group{}
	require.NoError(t, faker.FakeObject(&g))
	km := types.AccessKeyMetadata{}
	require.NoError(t, faker.FakeObject(&km))
	aup := types.AttachedPolicy{}
	require.NoError(t, faker.FakeObject(&aup))
	akl := iam.GetAccessKeyLastUsedOutput{}
	require.NoError(t, faker.FakeObject(&akl))

	sshPublicKey := types.SSHPublicKeyMetadata{}
	require.NoError(t, faker.FakeObject(&sshPublicKey))

	var tags []types.Tag
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListUsersOutput{
			Users: []types.User{u},
		}, nil)
	m.EXPECT().GetUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetUserOutput{
			User: &u,
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsForUserOutput{
			Groups: []types.Group{g},
		}, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []types.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []types.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&akl, nil)

	var l []string
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().ListUserPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListUserPoliciesOutput{
			PolicyNames: l,
		}, nil)

	p := iam.GetUserPolicyOutput{}
	require.NoError(t, faker.FakeObject(&p))
	document := "{\"test\": {\"t1\":1}}"
	p.PolicyDocument = &document
	m.EXPECT().GetUserPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	m.EXPECT().ListSSHPublicKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListSSHPublicKeysOutput{
			SSHPublicKeys: []types.SSHPublicKeyMetadata{sshPublicKey},
		}, nil)

	sc := types.SigningCertificate{}
	require.NoError(t, faker.FakeObject(&sc))
	m.EXPECT().ListSigningCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListSigningCertificatesOutput{
			Certificates: []types.SigningCertificate{sc},
		}, nil)

	m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.GenerateServiceLastAccessedDetailsOutput{JobId: aws.String("JobId")}, nil)

	lastAccessed := []types.ServiceLastAccessed{}
	require.NoError(t, faker.FakeObject(&lastAccessed))
	m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetServiceLastAccessedDetailsOutput{ServicesLastAccessed: lastAccessed, JobStatus: types.JobStatusTypeCompleted},
		nil,
	)

	return client.Services{
		Iam: m,
	}
}

func TestIamUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildIamUsers, client.TestOptions{})
}
