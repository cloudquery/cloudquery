package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func Test_doValidatePostgresVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		value      string
		minVersion string
		mockErr    error
		wantErr    error
	}{
		{
			"scan error",
			"",
			"10.0",
			errors.New("scan"),
			errors.New("error getting PostgreSQL version: scany: query one result row: scan"),
		},
		{
			"strange version output",
			"MSSQL",
			"10.0",
			nil,
			errors.New("error getting PostgreSQL version: failed to parse version: MSSQL"),
		},
		{
			"unparsable version",
			"PostgreSQL 10.a.1",
			"10.0",
			nil,
			errors.New("error getting PostgreSQL version: Malformed version: 10.a.1"),
		},
		{
			"lower than needed",
			"PostgreSQL 9.5 blah blah",
			"10.0",
			nil,
			errors.New("unsupported PostgreSQL version: 9.5.0. (should be >= 10.0.0)"),
		},
		{
			"equal",
			"PostgreSQL 10.0 blah blah",
			"10.0",
			nil,
			nil,
		},
		{
			"greater than needed",
			"PostgreSQL 12.5 blah blah",
			"10.0",
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			pgxRows := pgxpoolmock.NewRows([]string{"value"}).AddRow(tt.value).ToPgxRows()

			if tt.mockErr == nil {
				mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
			} else {
				mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, tt.mockErr)
			}

			want := version.Must(version.NewVersion(tt.minVersion))
			err := doValidatePostgresVersion(context.Background(), mockPool, want)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("wantErr is %v, returned error is %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
