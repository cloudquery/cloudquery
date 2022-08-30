package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	heroku "github.com/heroku/heroku-go/v5"
)

func TestParseNextRange(t *testing.T) {
	cases := []struct {
		give string
		want *heroku.ListRange
	}{
		{
			give: "id ..; max=10;",
			want: &heroku.ListRange{
				Field:      "id",
				Max:        10,
				Descending: false,
				FirstID:    "",
				LastID:     "",
			},
		},
		{
			give: "name from..to; order=desc,max=100;",
			want: &heroku.ListRange{
				Field:      "name",
				Max:        100,
				Descending: true,
				FirstID:    "from",
				LastID:     "to",
			},
		},
		{
			give: "id ]c46d9660-4da3-4cdc-9727-47a1d7e64834..; max=1000",
			want: &heroku.ListRange{
				Field:      "id",
				Max:        1000,
				Descending: false,
				FirstID:    "]c46d9660-4da3-4cdc-9727-47a1d7e64834",
				LastID:     "",
			},
		},
	}
	for _, tc := range cases {
		got, err := parseNextRange(tc.give)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("diff (-got, +want): %v", diff)
		}
	}
}
