package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProviderName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name             string
		args             args
		wantOrg          string
		wantProviderName string
		wantErr          bool
	}{
		{
			name:             "should return default org and name when given only name",
			args:             args{name: "test-name"},
			wantOrg:          DefaultOrganization,
			wantProviderName: "test-name",
			wantErr:          false,
		},
		{
			name:             "should return org and name when given org and name",
			args:             args{name: "test-org/test-name"},
			wantOrg:          "test-org",
			wantProviderName: "test-name",
			wantErr:          false,
		},
		{
			name:             "should error on invalid format",
			args:             args{name: "test-org/test-name/invalid"},
			wantOrg:          "",
			wantProviderName: "",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrg, gotProviderName, err := ParseProviderName(tt.args.name)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantOrg, gotOrg)
			assert.Equal(t, tt.wantProviderName, gotProviderName)
		})
	}
}

func TestProviderRepoName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return repo when give name",
			args: args{name: "test"},
			want: "cq-provider-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProviderRepoName(tt.args.name)
			assert.Equal(t, tt.want, got)
		})
	}
}
