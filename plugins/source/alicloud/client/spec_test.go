package client

import "testing"

import "github.com/cloudquery/codegen/jsonschema"

func TestSpecValidate(t *testing.T) {
	testCases := []struct {
		name string
		spec Spec
		err  bool
	}{
		{
			name: "valid spec",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "account1",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "key",
						SecretKey: "secret",
					},
				},
			},
		},
		{
			name: "valid spec with multiple accounts",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "account1",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "key",
						SecretKey: "secret",
					},
					{
						Name:      "account2",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "key",
						SecretKey: "secret",
					},
				},
			},
		},
		{
			name: "empty account name",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "key",
						SecretKey: "secret",
					},
				},
			},
			err: true,
		},
		{
			name: "empty regions",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "account1",
						Regions:   []string{},
						AccessKey: "key",
						SecretKey: "secret",
					},
				},
			},
			err: true,
		},
		{
			name: "empty key",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "account1",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "",
						SecretKey: "secret",
					},
				},
			},
			err: true,
		},
		{
			name: "empty secret",
			spec: Spec{
				Accounts: []AccountSpec{
					{
						Name:      "account1",
						Regions:   []string{"cn-hangzhou"},
						AccessKey: "key",
						SecretKey: "",
					},
				},
			},
			err: true,
		},
		{
			name: "empty spec",
			spec: Spec{},
			err:  true,
		},
		{
			name: "empty accounts",
			spec: Spec{
				Accounts: []AccountSpec{},
			},
			err: true,
		},
		{
			name: "null accounts",
			spec: Spec{
				Accounts: nil,
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.spec.Validate()
			if tc.err {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "valid spec",
			Spec: `{"accounts": [{"name": "account1", "regions": ["cn-hangzhou"], "access_key": "key", "secret_key": "secret"}]}`,
		},
		{
			Name: "missing account name",
			Spec: `{"accounts": [{"name": "", "regions": ["cn-hangzhou"], "access_key": "key", "secret_key": "secret"}]}`,
			Err:  true,
		},
		{
			Name: "missing access key",
			Spec: `{"accounts": [{"name": "account1", "regions": ["cn-hangzhou"], "access_key": "", "secret_key": "secret"}]}`,
			Err:  true,
		},
		{
			Name: "missing secret key",
			Spec: `{"accounts": [{"name": "account1", "regions": ["cn-hangzhou"], "access_key": "key", "secret_key": ""}]}`,
			Err:  true,
		},
		{
			Name: "missing regions",
			Spec: `{"accounts": [{"name": "account1", "regions": [], "access_key": "key", "secret_key": "secret"}]}`,
			Err:  true,
		},
		{
			Name: "empty spec",
			Spec: "{}",
			Err:  true,
		},
		{
			Name: "empty accounts",
			Spec: `{"accounts": []}`,
			Err:  true,
		},
		{
			Name: "null accounts",
			Spec: `{"accounts":null}`,
			Err:  true,
		},
	})
}
