package plugin

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newClient(t *testing.T) {
	t.Run("it should build an uninitialized client with no connection", func(t *testing.T) {
		options := plugin.NewClientOptions{NoConnection: true}

		client, err := newClient(context.TODO(), zerolog.Nop(), nil, options)
		assert.NoError(t, err)

		if assert.NotNil(t, client) {
			c := client.(*Client)
			assert.Nil(t, c.scheduler, "a client with no connection must not have a scheduler set")
			assert.Nil(t, c.client, "a client with no connection must not have a datadog client set")
		}
	})

	t.Run("it should return an error if the spec is not a valid JSON", func(t *testing.T) {
		invalidSpecFormat := []byte(`not a valid JSON object`)

		client, err := newClient(context.TODO(), zerolog.Nop(), invalidSpecFormat, plugin.NewClientOptions{})
		assert.ErrorContains(t, err, "invalid character")
		assert.Nil(t, client)
	})

	t.Run("it should return an error if the spec is not valid", func(t *testing.T) {
		specWithNoAccount := []byte(`{}`)

		client, err := newClient(context.TODO(), zerolog.Nop(), specWithNoAccount, plugin.NewClientOptions{})
		assert.ErrorContains(t, err, "no datadog accounts configured")
		assert.Nil(t, client)
	})

	t.Run("it should return an initialized client", func(t *testing.T) {
		specWithAccount := []byte(`{"accounts": [{ "name": "sample" }]}`)

		client, err := newClient(context.TODO(), zerolog.Nop(), specWithAccount, plugin.NewClientOptions{})
		assert.NoError(t, err)

		if assert.NotNil(t, client) {
			c := client.(*Client)
			assert.NotNil(t, c.scheduler, "a client with a connection must have a scheduler set")
			assert.NotNil(t, c.client, "a client with a connection must have a datadog client set")
		}
	})
}
