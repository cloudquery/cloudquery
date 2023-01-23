package shopify

import (
	"encoding/json"
	"strings"
)

type Tags []string

func (t *Tags) MarshalJSON() ([]byte, error) {
	if t == nil || len(*t) == 0 {
		return []byte("[]"), nil
	}

	return json.Marshal(strings.Join(*t, ", "))
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	var strTags string
	if err := json.Unmarshal(data, &strTags); err != nil {
		return err
	}

	if strTags == "" {
		tt := Tags{}
		*t = tt
		return nil
	}

	parts := strings.Split(strTags, ",")
	tt := make(Tags, 0, len(parts))
	for _, part := range parts {
		if s := strings.TrimSpace(part); s != "" {
			tt = append(tt, s)
		}
	}

	*t = tt
	return nil
}
