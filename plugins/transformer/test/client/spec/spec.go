package spec

type Spec struct {
	FailImmediately         bool `json:"fail_immediately"`
	FailAfterNSourceRecords int  `json:"fail_after_n_source_records"`
}

func (s *Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	return nil
}
