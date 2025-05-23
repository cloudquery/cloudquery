package secrets

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedaction(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
		env  []string
	}{
		{
			name: "replaces env var value with its name",
			env:  []string{"DB_PASS=foobar123"},
			msg:  "wrong password foobar123",
			want: "wrong password DB_PASS",
		},
		{
			name: "handles env var value with equals sign",
			env:  []string{"SECRET=user=foo"},
			msg:  "wrong password for user=foo",
			want: "wrong password for SECRET",
		},
		{
			name: "leaves original msg unchanged",
			env:  []string{},
			msg:  "wrong password foobar123",
			want: "wrong password foobar123",
		},
		{
			name: "handles env var with empty value",
			env:  []string{"DB_PASS="},
			msg:  "wrong password foobar123",
			want: "wrong password foobar123",
		},
		{
			name: "does not redact skipped env variables - AWS",
			env:  []string{"AWS_REGION=us-west-1"},
			msg:  "us-west-1: sample log",
			want: "us-west-1: sample log",
		},
		{
			name: "does not redact skipped env variables - HOME",
			env:  []string{"HOME=/root"},
			msg:  "sample log /root",
			want: "sample log /root",
		},
		{
			name: "does not redact skipped env variables - _CQ_TEAM_NAME",
			env:  []string{"_CQ_TEAM_NAME=team"},
			msg:  "wrong api key for team",
			want: "wrong api key for team",
		},
		{
			name: "does not redact variables with values of less than 4 characters",
			env:  []string{"DB_PASS=123"},
			msg:  "wrong password foobar123",
			want: "wrong password foobar123",
		},
		{
			name: "does redact CLOUDQUERY_API_KEY",
			env:  []string{"CLOUDQUERY_API_KEY=cqtk_test"},
			msg:  "wrong api key cqtk_test",
			want: "wrong api key CLOUDQUERY_API_KEY",
		},
	}
	for _, tt := range tests {
		redactor := NewSecretAwareRedactor()
		redactor.AddSecretEnv(tt.env)
		t.Run(tt.name, func(t *testing.T) {
			got := redactor.RedactStr(tt.msg)
			assert.Equal(t, tt.want, got)
		})
		t.Run(tt.name+" in log", func(t *testing.T) {
			out := &bytes.Buffer{}
			writer := NewSecretAwareWriter(out, redactor)
			_, _ = writer.Write([]byte(tt.msg))
			got, _ := io.ReadAll(out)
			assert.Equal(t, tt.want, string(got))
		})
	}
}
