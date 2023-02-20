package client

type Spec struct {
	Endpoint    string   `json:"endpoint,omitempty"`
	RegionCodes []string `json:"region_codes,omitempty"`
	OfferCodes  []string `json:"offer_codes,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Endpoint == "" {
		s.Endpoint = "https://pricing.us-east-1.amazonaws.com"
	}
	if len(s.RegionCodes) == 0 {
		s.RegionCodes = []string{"*"}
	}
	if len(s.OfferCodes) == 0 {
		s.OfferCodes = []string{"*"}
	}

}
