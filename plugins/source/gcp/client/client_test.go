package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUnion(t *testing.T) {
	assert.ElementsMatch(t, []string{"a", "b", "c"}, setUnion([]string{"a", "b"}, []string{"b", "c"}))
	assert.ElementsMatch(t, []string{"a", "b"}, setUnion([]string{"a", "b"}, []string{"a"}))
	assert.ElementsMatch(t, []string{"a", "b"}, setUnion([]string{"a", "b"}, []string{}))
}
