package models

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

type ImageScanWrapper struct {

	// The tag used for the image.
	ImageTag    *string
	ImageDigest *string

	*types.ImageScanFindings

	// The current state of the scan.
	*types.ImageScanStatus

	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string
}
