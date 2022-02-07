package client

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/go-version"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestMaybeCheckForUpdate(t *testing.T) {
	ctx := context.Background()
	const lastUpdateCheckPath = ".cq/last-update-check"
	viper.Set("data-dir", "./.cq")

	tests := []struct {
		name           string
		init           func(t *testing.T, fs afero.Afero)
		currentVersion string
		githubVersion  string
		githubError    error
		nowUnix        int64
		period         int64
		want           *version.Version
		wantErr        bool
		post           func(t *testing.T, fs afero.Afero)
	}{
		{
			name:           "development version",
			currentVersion: DevelopmentVersion,
			githubVersion:  "1.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        false,
			post: func(t *testing.T, fs afero.Afero) {
				ok, err := fs.Exists(lastUpdateCheckPath)
				require.Nil(t, err)
				require.False(t, ok)
			},
		},
		{
			name: "failed to read/write last-update-check",
			init: func(t *testing.T, fs afero.Afero) {
				require.Nil(t, fs.WriteFile(lastUpdateCheckPath+"/isdir", nil, 0o644))
			},
			currentVersion: "1.0.0",
			githubVersion:  "2.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        true,
		},
		{
			name: "last-update-check contains disable",
			init: func(t *testing.T, fs afero.Afero) {
				require.Nil(t, fs.WriteFile(lastUpdateCheckPath, []byte(" disabled  "), 0o644))
			},
			currentVersion: "1.0.0",
			githubVersion:  "2.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        false,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, " disabled  ", string(b)) // no change
			},
		},
		{
			name: "last-update-check has newer than current version",
			init: func(t *testing.T, fs afero.Afero) {
				require.Nil(t, fs.WriteFile(lastUpdateCheckPath, []byte("50 1.5.0"), 0o644))
			},
			currentVersion: "1.0.0",
			githubVersion:  "0.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           version.Must(version.NewSemver("1.5.0")),
			wantErr:        false,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, "50 1.5.0", string(b)) // no change
			},
		},
		{
			name: "last-update-check has the same version and period has not passed since last check",
			init: func(t *testing.T, fs afero.Afero) {
				require.Nil(t, fs.WriteFile(lastUpdateCheckPath, []byte("90 1.5.0"), 0o644))
			},
			currentVersion: "1.5.0",
			githubVersion:  "2.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        false,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, "90 1.5.0", string(b)) // no change
			},
		},
		{
			name:           "github reports an error",
			currentVersion: "1.0.0",
			githubVersion:  "2.0.0",
			githubError:    errors.New("fake"),
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        true,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, "89 0.0.0", string(b))
			},
		},
		{
			name:           "github returns unparsable version",
			currentVersion: "1.0.0",
			githubVersion:  "weirdTag",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           nil,
			wantErr:        true,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, "89 0.0.0", string(b))
			},
		},
		{
			name: "newer version is available on github",
			init: func(t *testing.T, fs afero.Afero) {
				require.Nil(t, fs.WriteFile(lastUpdateCheckPath, []byte("50 1.5.0"), 0o644))
			},
			currentVersion: "1.6.0",
			githubVersion:  "2.0.0",
			githubError:    nil,
			nowUnix:        100,
			period:         10,
			want:           version.Must(version.NewSemver("2.0.0")),
			wantErr:        false,
			post: func(t *testing.T, fs afero.Afero) {
				b, err := fs.ReadFile(lastUpdateCheckPath)
				require.Nil(t, err)
				require.Equal(t, "100 2.0.0", string(b))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveVersion := Version
			defer func() { Version = saveVersion }()
			Version = tt.currentVersion

			saveGetLatestRelease := getLatestRelease
			defer func() { getLatestRelease = saveGetLatestRelease }()
			getLatestRelease = func(ctx context.Context, client *http.Client, owner, repo string) (*github.RepositoryRelease, error) {
				return &github.RepositoryRelease{TagName: &tt.githubVersion}, tt.githubError
			}

			fs := afero.Afero{Fs: afero.NewMemMapFs()}
			if tt.init != nil {
				tt.init(t, fs)
			}
			got, err := MaybeCheckForUpdate(ctx, fs, tt.nowUnix, tt.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaybeCheckForUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("MaybeCheckForUpdate() = %v, want %v", got, tt.want)
			}
			if tt.post != nil {
				tt.post(t, fs)
			}
		})
	}
}
