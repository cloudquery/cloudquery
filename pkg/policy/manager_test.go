package policy

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_Load(t *testing.T) {
	cases := []struct {
		Name           string
		Policy         *Policy
		ExpectedPolicy *Policy
		ErrorOutput    string
	}{
		{
			Name: "load github policy",
			Policy: &Policy{
				Name:   "test",
				Source: "github.com/cloudquery-policies/test_policy",
			},
			ExpectedPolicy: &Policy{
				Name:        "test",
				Description: "this is a test policy",
				Doc:         "MAIN README",
				Policies: Policies{
					{
						Name:        "sub-policy",
						Description: "sub policy description",
						Doc:         "README FOR SUBPOLICY",
						Checks: []*Check{
							{
								Name:         "check",
								Description:  "test check",
								Doc:          "some doc md",
								ExpectOutput: true,
								Type:         AutomaticQuery,
								Query:        "SELECT 1;",
							},
						},
					},
				},
				Source: "github.com/cloudquery-policies/test_policy",
				meta: &Meta{
					Type:      "github",
					Version:   "",
					subPolicy: "",
					Directory: "cq/policies/manager/github.com/cloudquery-policies/test_policy",
				},
			},
		},
	}

	_ = os.RemoveAll(".cq/policies/manager")
	m := NewManager("./cq/policies/manager", nil, nil)
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			p, err := m.Load(context.Background(), tc.Policy)
			assert.Nil(t, err)
			assert.NotNil(t, p)
			p.meta.Directory = filepath.ToSlash(p.meta.Directory)
			assert.Equal(t, tc.ExpectedPolicy, p)
		})
	}

}
