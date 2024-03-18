package spec

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

// CloudQuery Azure source plugin retry options.
type RetryOptions struct {
	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L90).
	MaxRetries *int32 `json:"max_retries" jsonschema:"example=3"`

	// Disabled by default. Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L95).
	TryTimeoutSeconds *int `json:"try_timeout_seconds" jsonschema:"minimum=0,example=0"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L101).
	RetryDelaySeconds *int `json:"retry_delay_seconds" jsonschema:"minimum=0,example=4"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L106).
	MaxRetryDelaySeconds *int `json:"max_retry_delay_seconds" jsonschema:"minimum=0,example=60"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L118).
	//
	// The default of `null` uses the [default status codes](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L109).
	// An empty value disables retries for HTTP status codes.
	StatusCodes []int `json:"status_codes" jsonschema:"uniqueItems=true,example=408,example=500"`
}

func (r *RetryOptions) FillIn(options *policy.RetryOptions) {
	if r == nil || options == nil {
		return
	}

	if r.MaxRetries != nil {
		options.MaxRetries = *r.MaxRetries
	}
	if r.TryTimeoutSeconds != nil {
		options.TryTimeout = time.Duration(*r.TryTimeoutSeconds) * time.Second
	}
	if r.RetryDelaySeconds != nil {
		options.RetryDelay = time.Duration(*r.RetryDelaySeconds) * time.Second
	}
	if r.MaxRetryDelaySeconds != nil {
		options.MaxRetryDelay = time.Duration(*r.MaxRetryDelaySeconds) * time.Second
	}
	if r.StatusCodes != nil {
		options.StatusCodes = r.StatusCodes
	}
}
