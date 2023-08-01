package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	RegionCodes []string `json:"region_codes,omitempty"`
	OfferCodes  []string `json:"offer_codes,omitempty"`

	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
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
