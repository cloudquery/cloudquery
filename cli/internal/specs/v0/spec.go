package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

type Kind int

type Spec struct {
	Kind Kind `json:"kind"`
	Spec any  `json:"spec"`
}

const (
	KindSource Kind = iota
	KindDestination
)

func (k Kind) String() string {
	return [...]string{"source", "destination"}[k]
}

func (k Kind) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(k.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (k *Kind) UnmarshalJSON(data []byte) (err error) {
	var kind string
	if err := json.Unmarshal(data, &kind); err != nil {
		return err
	}
	if *k, err = KindFromString(kind); err != nil {
		return err
	}
	return nil
}

func KindFromString(s string) (Kind, error) {
	switch s {
	case "source":
		return KindSource, nil
	case "destination":
		return KindDestination, nil
	default:
		return KindSource, fmt.Errorf("unknown kind %s", s)
	}
}

func (s *Spec) UnmarshalJSON(data []byte) error {
	var t struct {
		Kind Kind `json:"kind"`
		Spec any  `json:"spec"`
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	if err := dec.Decode(&t); err != nil {
		return err
	}
	s.Kind = t.Kind
	switch s.Kind {
	case KindSource:
		s.Spec = new(Source)
	case KindDestination:
		s.Spec = new(Destination)
	default:
		return fmt.Errorf("unknown kind %s", s.Kind)
	}
	b, err := json.Marshal(t.Spec)
	if err != nil {
		return err
	}
	dec = json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.DisallowUnknownFields()
	return dec.Decode(s.Spec)
}

func UnmarshalJSONStrict(b []byte, out any) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	return dec.Decode(out)
}

func SpecUnmarshalYamlStrict(b []byte, spec *Spec) error {
	jb, err := yaml.YAMLToJSON(b)
	if err != nil {
		return fmt.Errorf("failed to convert yaml to json: %w", err)
	}
	dec := json.NewDecoder(bytes.NewReader(jb))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	if err := dec.Decode(spec); err != nil {
		return fmt.Errorf("failed to decode spec: %w", err)
	}
	return nil
}
