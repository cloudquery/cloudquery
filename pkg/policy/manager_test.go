package policy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_Load(t *testing.T) {
	cases := []struct {
		Name        string
		Policy      *Policy
		ErrorOutput string
	}{
		{
			Name: "load github policy",
			Policy: &Policy{
				Name:   "aws",
				Source: "github.com/cloudquery-policies/aws?ref=policy_v3",
			},
		},
	}

	m := NewManager(".", nil, nil)
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			p, err := m.Load(context.Background(), tc.Policy)
			assert.Nil(t, err)
			assert.NotNil(t, p)
		})
	}

}
