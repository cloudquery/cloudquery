package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Strip(t *testing.T) {
	assert.Equal(t, strip("\\u001b[31m❌\\u001b[0m failed to sync provider k8s@latest"), "failed to sync provider k8s@latest")
	assert.Equal(t, strip("\\u001b[31m\\u001b[0m failed to sync provider k8s@latest."), "failed to sync provider k8s@latest")
}
