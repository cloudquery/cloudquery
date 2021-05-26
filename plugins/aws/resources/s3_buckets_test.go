package resources

import (
	"math/rand"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildS3Buckets(t *testing.T, ctrl *gomock.Controller) client.Services {
	mgr := mocks.NewMockS3ManagerClient(ctrl)
	m := mocks.NewMockS3Client(ctrl)
	b := s3Types.Bucket{}
	err := faker.FakeData(&b)
	if err != nil {
		t.Fatal(err)
	}
	bloc := s3.GetBucketLocationOutput{}
	err = faker.FakeData(&bloc)
	if err != nil {
		t.Fatal(err)
	}
	blog := s3.GetBucketLoggingOutput{}
	err = faker.FakeData(&blog)
	if err != nil {
		t.Fatal(err)
	}
	bpol := s3.GetBucketPolicyOutput{}
	err = faker.FakeData(&bpol)
	if err != nil {
		t.Fatal(err)
	}
	jsonDoc := `{"stuff": 3}`
	bpol.Policy = &jsonDoc
	bver := s3.GetBucketVersioningOutput{}
	err = faker.FakeData(&bver)
	if err != nil {
		t.Fatal(err)
	}
	bgrant := s3Types.Grant{}
	err = faker.FakeData(&bgrant)
	if err != nil {
		t.Fatal(err)
	}
	bcors := s3Types.CORSRule{}
	err = faker.FakeData(&bcors)
	if err != nil {
		t.Fatal(err)
	}
	bencryption := s3.GetBucketEncryptionOutput{}
	err = faker.FakeData(&bencryption)
	if err != nil {
		t.Fatal(err)
	}

	bpba := s3.GetPublicAccessBlockOutput{}
	err = faker.FakeData(&bpba)
	if err != nil {
		t.Fatal(err)
	}
	btag := s3.GetBucketTaggingOutput{}
	err = faker.FakeData(&btag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any()).Return(
		&s3.ListBucketsOutput{
			Buckets: []s3Types.Bucket{b},
		}, nil)
	m.EXPECT().GetBucketLogging(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&blog, nil)
	m.EXPECT().GetBucketPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&bpol, nil)
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

	// bucket replication struct has interfaces and faker doesn't work well with it so we will build it manually
	sourceSelectionCriteria := s3Types.SourceSelectionCriteria{}
	if err := faker.FakeData(&sourceSelectionCriteria); err != nil {
		t.Fatal(err)
	}
	replicationDest := s3Types.Destination{}
	if err := faker.FakeData(&replicationDest); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBucketReplication(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketReplicationOutput{
			ReplicationConfiguration: &s3Types.ReplicationConfiguration{
				Role: aws.String(faker.Name()),
				Rules: []s3Types.ReplicationRule{{
					Destination:               &replicationDest,
					Status:                    s3Types.ReplicationRuleStatusDisabled,
					DeleteMarkerReplication:   &s3Types.DeleteMarkerReplication{Status: s3Types.DeleteMarkerReplicationStatusEnabled},
					ExistingObjectReplication: &s3Types.ExistingObjectReplication{Status: s3Types.ExistingObjectReplicationStatusEnabled},
					Filter:                    &s3Types.ReplicationRuleFilterMemberPrefix{Value: "blabla"},
					ID:                        aws.String(faker.Name()),
					Prefix:                    aws.String(faker.Name()),
					Priority:                  5,
					SourceSelectionCriteria:   &sourceSelectionCriteria,
				}},
			}}, nil)
	m.EXPECT().GetBucketTagging(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&btag, nil)
	randomTime, _ := time.Parse(time.RFC3339, faker.Timestamp())
	m.EXPECT().GetBucketLifecycleConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3.GetBucketLifecycleConfigurationOutput{Rules: []s3Types.LifecycleRule{{
			Status:                         s3Types.ExpirationStatusEnabled,
			AbortIncompleteMultipartUpload: &s3Types.AbortIncompleteMultipartUpload{DaysAfterInitiation: rand.Int31()},
			Expiration:                     &s3Types.LifecycleExpiration{Date: &randomTime, Days: 3, ExpiredObjectDeleteMarker: false},
			Filter:                         &s3Types.LifecycleRuleFilterMemberPrefix{Value: "blabla"},
			ID:                             aws.String(faker.Name()),
			NoncurrentVersionExpiration:    &s3Types.NoncurrentVersionExpiration{NoncurrentDays: 33},
			NoncurrentVersionTransitions:   []s3Types.NoncurrentVersionTransition{{NoncurrentDays: 5, StorageClass: s3Types.TransitionStorageClassDeepArchive}},
			Prefix:                         aws.String(faker.Name()),
			Transitions: []s3Types.Transition{{
				Date:         &randomTime,
				Days:         15,
				StorageClass: s3Types.TransitionStorageClassOnezoneIa,
			}},
		}}}, nil)
	mgr.EXPECT().GetBucketRegion(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		"us-east-1", nil)
	return client.Services{
		S3:        m,
		S3Manager: mgr,
	}
}

func TestS3Buckets(t *testing.T) {
	awsTestHelper(t, S3Buckets(), buildS3Buckets)
}
