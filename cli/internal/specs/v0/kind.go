package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Kind int

const (
	KindSource Kind = iota
	KindDestination
	KindTransformer
)

var (
	AllKinds = [...]string{
		KindSource:      "source",
		KindDestination: "destination",
		KindTransformer: "transformer",
	}
)

func (k Kind) String() string {
	return AllKinds[k]
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

func (Kind) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllKinds))
	for i, k := range AllKinds {
		sc.Enum[i] = k
	}
}

func KindFromString(s string) (Kind, error) {
	for k, str := range AllKinds {
		if s == str {
			return Kind(k), nil
		}
	}
	return KindSource, fmt.Errorf("unknown kind %s", s)
}
