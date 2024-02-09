package models

import "github.com/aws/aws-sdk-go-v2/service/guardduty"

type DetectorWrapper struct {
	*guardduty.GetDetectorOutput
	Id string
}

type IPSetWrapper struct {
	*guardduty.GetIPSetOutput
	Id string
}

type ThreatIntelSetWrapper struct {
	*guardduty.GetThreatIntelSetOutput
	Id string
}
