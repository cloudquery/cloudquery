package client

type Spec struct {
	Directory string `json:"directory,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Directory == "" {
		s.Directory = "./cq_csv_output"
	}
}
