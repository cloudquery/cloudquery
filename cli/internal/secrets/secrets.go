package secrets

import (
	"bytes"
	"io"
	"slices"
	"strings"
)

var skippedEnvironmentVariables = []string{
	// injected by Kubernetes
	"JOB_COMPLETION_INDEX",

	// injected by the CloudQuery CLI
	"CQ_CLOUD",

	// injected by EKS, do not contain any sensitive information regardless
	"AWS_STS_REGIONAL_ENDPOINTS", "AWS_DEFAULT_REGION", "AWS_REGION",
	"AWS_WEB_IDENTITY_TOKEN_FILE", "AWS_ROLE_ARN", "AWS_ROLE_SESSION_NAME",
}

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
		if len(parts) != 2 || len(parts[1]) == 0 {
			continue
		}
		if slices.Contains(skippedEnvironmentVariables, parts[0]) {
			continue
		}
		s.secrets[parts[1]] = parts[0]
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
