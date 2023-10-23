package spec

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

// CloudQuery Azure source plugin retry options.
type RetryOptions struct {
	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L90).
	MaxRetries *int32 `json:"max_retries"`

	// Disabled by default. Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L95).
	TryTimeout *configtype.Duration `json:"try_timeout"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L101).
	RetryDelay *configtype.Duration `json:"retry_delay"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L106).
	MaxRetryDelay *configtype.Duration `json:"max_retry_delay"`

	// Described in the
	// [Azure Go SDK](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L118).
	StatusCodes []int `json:"status_codes" jsonschema:"uniqueItems=true"`
}

func (r *RetryOptions) FillIn(options *policy.RetryOptions) {
	if r == nil || options == nil {
		return
	}

	if r.MaxRetries != nil {
		options.MaxRetries = *r.MaxRetries
	}
	if r.TryTimeout != nil {
		options.TryTimeout = r.TryTimeout.Duration()
	}
	if r.RetryDelay != nil {
		options.RetryDelay = r.RetryDelay.Duration()
	}
	if r.MaxRetryDelay != nil {
		options.MaxRetryDelay = r.MaxRetryDelay.Duration()
	}
	if r.StatusCodes != nil {
		options.StatusCodes = r.StatusCodes
	}
}
