package spec

type Spec struct {
	FailImmediately         bool `json:"fail_immediately"`
	FailAfterNSourceRecords int  `json:"fail_after_n_source_records"`
	ExitImmediately         bool `json:"exit_immediately"`
}

func (*Spec) SetDefaults() {
}

func (*Spec) Validate() error {
	return nil
}
