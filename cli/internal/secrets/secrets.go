package secrets

import (
	"bytes"
	"io"
	"strings"
)

type SecretAwareRedactor struct {
	secrets map[string]string
}

func NewSecretAwareRedactor() *SecretAwareRedactor {
	return &SecretAwareRedactor{secrets: make(map[string]string)}
}

func (s *SecretAwareRedactor) RedactStr(msg string) string {
	return string(s.RedactBytes([]byte(msg)))
}

func (s *SecretAwareRedactor) RedactBytes(msg []byte) []byte {
	for v, k := range s.secrets {
		msg = bytes.ReplaceAll(msg, []byte(v), []byte(k))
	}
	return msg
}

func (s *SecretAwareRedactor) AddSecretEnv(env []string) {
	for _, v := range env {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 && len(parts[1]) > 0 {
			s.secrets[parts[1]] = parts[0]
		}
	}
}

type SecretAwareWriter struct {
	out      io.Writer
	redactor *SecretAwareRedactor
}

func NewSecretAwareWriter(out io.Writer, redactor *SecretAwareRedactor) *SecretAwareWriter {
	return &SecretAwareWriter{out: out, redactor: redactor}
}

func (s SecretAwareWriter) Write(p []byte) (n int, err error) {
	return s.out.Write(s.redactor.RedactBytes(p))
}
