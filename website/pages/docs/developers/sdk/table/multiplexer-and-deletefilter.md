import { Callout } from 'nextra-theme-docs'

# Multiplexer and DeleteFilter

`Multiplex` and `DeleteFilter` options are defined in the [schema.Table](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) entry, per resource type.

## Multiplexer

A multiplexer receives a client and returns an array of clients to execute in parallel:

```go
func(meta ClientMeta) []ClientMeta
```

It's used to parallelize the same resource across multiple accounts or service regions. A multiplexer usually makes copies of the given client, and sets a certain property to a different value on each copy.

For instance this one from the [AWS Provider](https://github.com/cloudquery/cq-provider-aws/blob/main/client/multiplexers.go) iterates every account and re-creates the client multiple times, one client per account:

```go
func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		l = append(l, client.withAccountID(accountID))
	}
	return l
}
```

## DeleteFilter

DeleteFilter defines how to remove a certain type of resource from the database, by returning a list of key/value pairs to match when truncating that resource from the database.

Again, this one from the [AWS Provider](https://github.com/cloudquery/cq-provider-aws/blob/main/client/filters.go) returns an `account_id=<THE_ACCOUNT_ID>` pair, by reading the Account ID from the given `*Client`.

```go
func DeleteAccountFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{"account_id", client.AccountID}
}
```

<Callout type="info">

DeleteFilter is always a _mirror_ of your Multiplexer.

</Callout>
