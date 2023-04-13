package client

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSpec_Validate(t *testing.T) {
	cases := []struct {
		name string
		spec Spec
		want error
	}{
		{
			name: "empty",
			spec: Spec{},
			want: errors.New("missing API key"),
		},
		{
			name: "valid",
			spec: Spec{
				APIKey: "test",
			},
		},
		{
			name: "period with to/from",
			spec: Spec{
				APIKey: "test",
				TableOptions: &TableOptions{
					SnykReportingIssues: SnykReportingIssuesOptions{
						From:   "2020-01-01",
						To:     "2020-01-01",
						Period: "1d",
					},
				},
			},
			want: errors.New("cannot use period with to/from"),
		},
		{
			name: "to without from",
			spec: Spec{
				APIKey: "test",
				TableOptions: &TableOptions{
					SnykReportingIssues: SnykReportingIssuesOptions{
						To:     "2020-01-01",
						Period: "1d",
					},
				},
			},
			want: errors.New("cannot use to without from"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.spec.Validate()
			if tc.want != nil {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.want.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSnykReportingIssuesOptions(t *testing.T) {
	cases := []struct {
		name     string
		opts     SnykReportingIssuesOptions
		wantTo   time.Time
		wantFrom time.Time
	}{
		{
			name:     "default",
			opts:     SnykReportingIssuesOptions{},
			wantTo:   time.Now().Truncate(24 * time.Hour),
			wantFrom: time.Now().Add(-parseDuration(defaultPeriod)).Truncate(24 * time.Hour),
		},
		{
			name: "to/from",
			opts: SnykReportingIssuesOptions{
				From: "2020-01-01",
				To:   "2020-01-02",
			},
			wantFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "period 7d",
			opts: SnykReportingIssuesOptions{
				Period: "7d",
			},
			wantFrom: time.Now().Add(-time.Hour * 24 * 7).Truncate(24 * time.Hour),
			wantTo:   time.Now().Truncate(24 * time.Hour),
		},
		{
			name: "period 2w",
			opts: SnykReportingIssuesOptions{
				Period: "2w",
			},
			wantFrom: time.Now().Add(-time.Hour * 24 * 14).Truncate(24 * time.Hour),
			wantTo:   time.Now().Truncate(24 * time.Hour),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			from, to := tc.opts.FromTime(), tc.opts.ToTime()
			require.Equal(t, tc.wantTo, to)
			require.Equal(t, tc.wantFrom, from)
		})
	}
}
