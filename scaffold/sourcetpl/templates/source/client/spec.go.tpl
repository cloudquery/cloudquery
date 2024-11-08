package client

type Spec struct {
	// plugin spec goes here
}

func (s *Spec) SetDefaults() {
	// set defaults for the spec
}

func (s *Spec) Validate() error {
	// validate the spec
	return nil
}
