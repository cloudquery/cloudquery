package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

type ConformancePackComplianceWrapper struct {
	// Fields from types.ConformancePackRuleCompliance:

	// Compliance of the Config rule. The allowed values are COMPLIANT, NON_COMPLIANT,
	// and INSUFFICIENT_DATA.
	ComplianceType types.ConformancePackComplianceType

	// Name of the Config rule.
	ConfigRuleName *string

	// Controls for the conformance pack. A control is a process to prevent or detect
	// problems while meeting objectives. A control can align with a specific
	// compliance regime or map to internal controls defined by an organization.
	Controls []string

	// Fields from types.ConformancePackEvaluationResult:

	// The time when Config rule evaluated Amazon Web Services resource.
	//
	// This member is required.
	ConfigRuleInvokedTime *time.Time

	// Uniquely identifies an evaluation result.
	//
	// This member is required.
	EvaluationResultIdentifier *types.EvaluationResultIdentifier

	// The time when Config recorded the evaluation result.
	//
	// This member is required.
	ResultRecordedTime *time.Time

	// Supplementary information about how the evaluation determined the compliance.
	Annotation *string
}

type ConfigurationRecorderWrapper struct {
	types.ConfigurationRecorder
	StatusLastErrorCode        *string
	StatusLastErrorMessage     *string
	StatusLastStartTime        *time.Time
	StatusLastStatus           types.RecorderStatus
	StatusLastStatusChangeTime *time.Time
	StatusLastStopTime         *time.Time
	StatusRecording            bool
}
