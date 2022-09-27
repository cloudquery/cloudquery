package models

import "github.com/aws/aws-sdk-go-v2/service/guardduty"

type DetectorWrapper struct {
	*guardduty.GetDetectorOutput
	Id string
}
