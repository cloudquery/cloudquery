package client

import _ "embed"

type Spec struct {
	// List of region codes to sync. If empty, everything will be synced.
	RegionCodes []string `json:"region_codes,omitempty" jsonschema:"minLength=1"`

	// List of offer codes to sync. If empty, everything will be synced.
	OfferCodes []string `json:"offer_codes,omitempty" jsonschema:"minLength=1"`

	// Concurrency setting for the CloudQuery scheduler
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) SetDefaults() {
	if len(s.RegionCodes) == 0 {
		s.RegionCodes = []string{"*"}
	}
	if len(s.OfferCodes) == 0 {
		s.OfferCodes = []string{"*"}
	}

	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (*Spec) Validate() error {
	return nil
}

//go:embed schema.json
var JSONSchema string
