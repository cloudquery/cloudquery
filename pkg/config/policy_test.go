package config

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicyWrapper_Normalize(t *testing.T) {
	// this test asserts the fact that Normalize is called on all inner policies
	tests := []struct {
		before PolicyWrapper
		after  PolicyWrapper
	}{
		{
			PolicyWrapper{Policies: []*Policy{
				{
					Queries: []*Query{
						{Type: ""},
					},
				},
				{
					Queries: []*Query{
						{Type: ""},
					},
				},
			}},
			PolicyWrapper{Policies: []*Policy{
				{
					Queries: []*Query{
						{Type: AutomaticQuery},
					},
				},
				{
					Queries: []*Query{
						{Type: AutomaticQuery},
					},
				},
			}},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.before.Normalize()
			assert.True(t, reflect.DeepEqual(tt.before, tt.after))
		})
	}
}

func TestPolicyWrapper_Validate(t *testing.T) {
	// this test asserts the fact that Validate is called on all inner policies
	tests := []struct {
		name    string
		pw      PolicyWrapper
		wantErr bool
	}{
		{
			"first policy fails",
			PolicyWrapper{Policies: []*Policy{
				{
					Queries: []*Query{
						{Type: ""},
					},
				},
				{
					Queries: []*Query{
						{Type: ManualQuery},
					},
				},
			}},
			true,
		},
		{
			"second policy fails",
			PolicyWrapper{Policies: []*Policy{
				{
					Queries: []*Query{
						{Type: AutomaticQuery},
					},
				},
				{
					Queries: []*Query{
						{Type: "wrong"},
					},
				},
			}},
			true,
		},
		{
			"OK",
			PolicyWrapper{Policies: []*Policy{
				{
					Queries: []*Query{
						{Type: AutomaticQuery},
					},
				},
				{
					Queries: []*Query{
						{Type: ManualQuery},
					},
				},
			}},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if err := tt.pw.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PolicyWrapper.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPolicy_Normalize(t *testing.T) {
	tests := []struct {
		name   string
		before Policy
		after  Policy
	}{
		{
			"Normalize is called on inner Queries",
			Policy{
				Queries: []*Query{
					{Type: ""},
				},
			},
			Policy{
				Queries: []*Query{
					{Type: AutomaticQuery},
				},
			},
		},
		{
			"Normalize is called on inner Policies",
			Policy{
				Policies: []*Policy{
					{
						Queries: []*Query{
							{Type: ""},
						},
					},
				},
			},
			Policy{
				Policies: []*Policy{
					{
						Queries: []*Query{
							{Type: AutomaticQuery},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before.Normalize()
			assert.True(t, reflect.DeepEqual(tt.before, tt.after))
		})
	}
}

func TestPolicy_Validate(t *testing.T) {
	tests := []struct {
		name    string
		p       Policy
		wantErr bool
	}{
		{
			"Validate is called on inner queries",
			Policy{
				Queries: []*Query{
					{Type: ""},
				},
			},
			true,
		},
		{
			"Validate is called on inner policies",
			Policy{
				Policies: []*Policy{
					{
						Queries: []*Query{
							{Type: ""},
						},
					},
				},
			},
			true,
		},
		{
			"Happy case",
			Policy{
				Queries: []*Query{
					{Type: AutomaticQuery},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Query.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuery_Normalize(t *testing.T) {
	tests := []struct {
		before Query
		after  Query
	}{
		{
			Query{Query: "ABC", Type: AutomaticQuery},
			Query{Query: "ABC", Type: AutomaticQuery},
		},
		{
			Query{Query: "ABC", Type: ManualQuery},
			Query{Query: "ABC", Type: ManualQuery},
		},
		{
			Query{Query: "ABC", Type: ""},
			Query{Query: "ABC", Type: AutomaticQuery},
		},
		{
			Query{Query: "ABC", Type: "unexpected"},
			Query{Query: "ABC", Type: "unexpected"},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.before.Normalize()
			assert.True(t, reflect.DeepEqual(tt.before, tt.after))
		})
	}
}

func TestQuery_Validate(t *testing.T) {
	tests := []struct {
		q       Query
		wantErr bool
	}{
		{Query{Query: "ABC", Type: ManualQuery}, false},
		{Query{Query: "ABC", Type: AutomaticQuery}, false},
		{Query{Query: "ABC", Type: "wrong"}, true},
		{Query{Query: "ABC", Type: ""}, true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if err := tt.q.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Query.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
