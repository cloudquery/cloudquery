package client

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/assert"
)

func TestLowercaseID(t *testing.T) {
	type test1 struct {
		ID string `json:"id,omitempty"`
	}
	assert.Equal(t, &test1{"id"}, lowercaseID(&test1{"ID"}))

	type test2 struct {
		ID *string `json:"id,omitempty"`
	}
	assert.Equal(t, &test2{to.Ptr("id")}, lowercaseID(&test2{to.Ptr("Id")}))

	type test3 struct {
		ID int64 `json:"id,omitempty"`
	}
	assert.Equal(t, &test3{123}, lowercaseID(&test3{123}))
}
