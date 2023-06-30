package client

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4 // 10MB
)

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	BatchSize        int    `json:"batch_size,omitempty"`
	BatchSizeBytes   int    `json:"batch_size_bytes,omitempty"`
	Debug            bool   `json:"debug,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
}
