package client

type Spec struct {
	ItemConcurrency int `json:"item_concurrency"`
}

func (s *Spec) SetDefaults() {
	if s.ItemConcurrency == 0 {
		s.ItemConcurrency = 100 // Default to loading 100 concurrent items
	}
}
