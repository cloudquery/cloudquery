package spec

type Spec struct {
	Tables []string `json:"tables"`
}

func (s *Spec) SetDefaults() {
	if s.Tables == nil {
		s.Tables = []string{"*"}
	}
}

func (*Spec) Validate() error {
	return nil
}
