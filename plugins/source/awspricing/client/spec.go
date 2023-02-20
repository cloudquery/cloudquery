package client

type Spec struct {
	RegionCodes []string `json:"region_codes,omitempty"`
	OfferCodes  []string `json:"offer_codes,omitempty"`
}

func (s *Spec) SetDefaults() {
	if len(s.RegionCodes) == 0 {
		s.RegionCodes = []string{"*"}
	}
	if len(s.OfferCodes) == 0 {
		s.OfferCodes = []string{"*"}
	}
}
