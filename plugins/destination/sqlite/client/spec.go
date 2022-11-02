package client

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	BatchSize        int    `json:"batch_size,omitempty"`
}

const defaultBatchSize = 1000

func (s *Spec) SetDefaults() {
	if s.BatchSize <= 0 {
		s.BatchSize = defaultBatchSize
	}
}
