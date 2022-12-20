package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

type PublicAccessBlockConfigurationWrapper struct {
	s3controlTypes.PublicAccessBlockConfiguration
	ConfigExists bool
}

type WrappedBucket struct {
	// CreationDate and Name are from types.Bucket:

	// Date the bucket was created. This date can change when making changes to your
	// bucket, such as editing its bucket policy.
	CreationDate *time.Time
	// The name of the bucket.
	Name *string

	// Fields obtained from other SDK calls:

	ReplicationRole       *string
	ReplicationRules      []types.ReplicationRule
	Region                string
	LoggingTargetBucket   *string
	LoggingTargetPrefix   *string
	Policy                map[string]any
	VersioningStatus      types.BucketVersioningStatus
	VersioningMfaDelete   types.MFADeleteStatus
	BlockPublicAcls       bool
	BlockPublicPolicy     bool
	IgnorePublicAcls      bool
	RestrictPublicBuckets bool
	Tags                  map[string]*string
	OwnershipControls     []string
}
