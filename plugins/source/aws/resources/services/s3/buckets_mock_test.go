package s3

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildS3Buckets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3Client(ctrl)
	b := s3Types.Bucket{}
	require.NoError(t, faker.FakeObject(&b))
	bloc := s3.GetBucketLocationOutput{}
	require.NoError(t, faker.FakeObject(&bloc))
	blog := s3.GetBucketLoggingOutput{}
	require.NoError(t, faker.FakeObject(&blog))
	bpol := s3.GetBucketPolicyOutput{}
	require.NoError(t, faker.FakeObject(&bpol))
	bpols := s3.GetBucketPolicyStatusOutput{}
	require.NoError(t, faker.FakeObject(&bpols))
	jsonDoc := `{"stuff": 3}`
	bpol.Policy = &jsonDoc
	bver := s3.GetBucketVersioningOutput{}
	require.NoError(t, faker.FakeObject(&bver))
	bgrant := s3Types.Grant{}
	require.NoError(t, faker.FakeObject(&bgrant))
	// set up properly
	bgrant.Grantee.EmailAddress = nil
	bgrant.Grantee.ID = nil
	bgrant.Grantee.Type = s3Types.TypeGroup

	bcors := s3Types.CORSRule{}
	require.NoError(t, faker.FakeObject(&bcors))
	bencryption := s3.GetBucketEncryptionOutput{}
	require.NoError(t, faker.FakeObject(&bencryption))

	bpba := s3.GetPublicAccessBlockOutput{}
	require.NoError(t, faker.FakeObject(&bpba))
	btag := s3.GetBucketTaggingOutput{}
	require.NoError(t, faker.FakeObject(&btag))
	bownershipcontrols := s3.GetBucketOwnershipControlsOutput{}
	require.NoError(t, faker.FakeObject(&bownershipcontrols))

	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.ListBucketsOutput{
			Buckets: []s3Types.Bucket{b},
		}, nil)
	m.EXPECT().GetBucketLogging(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&blog, nil)
	m.EXPECT().GetBucketPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bpol, nil)
	m.EXPECT().GetBucketPolicyStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bpols, nil)
	m.EXPECT().GetBucketVersioning(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bver, nil)
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketAclOutput{
			Grants: []s3Types.Grant{bgrant},
		}, nil)
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketCorsOutput{
			CORSRules: []s3Types.CORSRule{bcors},
		}, nil)
	m.EXPECT().GetBucketEncryption(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bencryption, nil)

	m.EXPECT().GetPublicAccessBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bpba, nil)
	m.EXPECT().GetBucketOwnershipControls(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bownershipcontrols, nil)

	ro := s3.GetBucketReplicationOutput{}
	require.NoError(t, faker.FakeObject(&ro))

	m.EXPECT().GetBucketReplication(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ro, nil)
	m.EXPECT().GetBucketTagging(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&btag, nil)
	tt := s3Types.Transition{}
	require.NoError(t, faker.FakeObject(&tt))

	glco := s3.GetBucketLifecycleConfigurationOutput{}
	require.NoError(t, faker.FakeObject(&glco))

	m.EXPECT().GetBucketLifecycleConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&glco, nil)
	m.EXPECT().GetBucketLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bloc, nil)

	websiteOutput := s3.GetBucketWebsiteOutput{}
	require.NoError(t, faker.FakeObject(&websiteOutput))

	m.EXPECT().GetBucketWebsite(gomock.Any(), gomock.Any(), gomock.Any()).Return(&websiteOutput, nil)

	return client.Services{
		S3: m,
	}
}

func TestS3Buckets(t *testing.T) {
	client.AwsMockTestHelper(t, Buckets(), buildS3Buckets, client.TestOptions{})
}
