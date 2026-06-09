package plugin

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/services"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	cases := []struct {
		name      string
		spec      string
		probeErr  error
		setupMock bool
		wantCode  string
	}{
		{name: "ok", spec: `{}`, setupMock: true},
		{name: "invalid spec", spec: `{"item_concurrency":"nope"}`, wantCode: codeInvalidSpec},
		{name: "invalid spec validate", spec: `{"item_concurrency":2000}`, wantCode: codeInvalidSpec},
		{name: "connection failed", spec: `{}`, setupMock: true, probeErr: errors.New("boom"), wantCode: codeConnectionFailed},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mocks.NewMockHackernewsClient(ctrl)
			if tc.setupMock {
				m.EXPECT().MaxItemID(gomock.Any()).Return(123, tc.probeErr)
			}
			tester := connectionTester(func() services.HackernewsClient { return m })
			err := tester(context.Background(), zerolog.Nop(), []byte(tc.spec))
			if tc.wantCode == "" {
				require.NoError(t, err)
				return
			}
			var connErr *plugin.TestConnError
			require.ErrorAs(t, err, &connErr)
			require.Equal(t, tc.wantCode, connErr.Code)
		})
	}
}
