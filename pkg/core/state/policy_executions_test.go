package state

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_CreatePolicyExecution(t *testing.T) {
	c, closer := initStateClient(t)
	defer closer()
	pe := &PolicyExecution{
		Scheme:     "foo",
		Location:   "bar",
		PolicyName: "baz",
		Selector:   "123",
		Sha256Hash: "321",
		Version:    "1.2.3",
	}

	data, err := c.db.Query(context.Background(), "SELECT * FROM cloudquery.policy_executions")
	assert.NoError(t, err)
	countBefore := 0
	for data.Next() {
		countBefore += 1
	}
	assert.NoError(t, data.Err())

	_, err = c.CreatePolicyExecution(context.Background(), pe)
	assert.NoError(t, err)

	data, err = c.db.Query(context.Background(), "SELECT * FROM cloudquery.policy_executions")
	assert.NoError(t, err)
	countAfter := 0
	for data.Next() {
		countAfter += 1
	}
	assert.NoError(t, data.Err())
	assert.Equal(t, countBefore+1, countAfter)
}

func initStateClient(t *testing.T) (*Client, func()) {
	sta, err := NewClient(context.Background(), testDBConnection)
	assert.NoError(t, err)
	return sta, sta.Close
}
