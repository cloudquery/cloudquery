import { Callout } from 'nextra-theme-docs'

# Fetch Resolvers

Fetch resolvers are functions to fetch resources from the source (a cloud provider for example).

## TableResolver

This is the main type of fetch resolver:

```go
type TableResolver func(ctx context.Context, meta ClientMeta, parent *Resource, res chan interface{}) error
```

`TableResolver` allows you to access the cloud resource using the given passed `*Client` and fetch all resources of that type. Finally, you send the fetch items into the passed `res` channel argument.
The `TableResolver` is flexible allowing you to define your own pagination logic or any other logic for that matter, and pass the results to the channel.

<Callout type="info">

The collector in the SDK is slice-aware, so if you have a slice of resources, you can just push the slice as a whole, without iterating.

</Callout>

Here's an example from the [provider template](https://github.com/cloudquery/cq-provider-template/blob/main/resources/services/demo/resource.go):

```go
func fetchDemoResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	_ = c
	// Fetch using the third party client and put the result in res
	// res <- c.ThirdPartyClient.getData()
	return nil
}
```

## PostResourceResolver (RowResolver)

This optional resolver is called after all columns have been resolved, and before resource is inserted to database. `PostResourceResolver` is the name in the `schema.Table` struct.

```go
type RowResolver func(ctx context.Context, meta ClientMeta, resource *Resource) error
```

Here's an example from the AWS Provider's [SNS Topics resource](https://github.com/cloudquery/cq-provider-aws/blob/006d963/resources/sns_topics.go), getting SNS topic attributes in one go and setting various fields:

```go
func resolveTopicAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	topic, ok := resource.Item.(types.Topic)

	// ...

	output, err := svc.GetTopicAttributes(ctx, &params, func(o *sns.Options) {
		o.Region = c.Region
	})

	// ...

	// Set all attributes
	if err := resource.Set("subscriptions_confirmed", cast.ToInt(output.Attributes["SubscriptionsConfirmed"])); err != nil {
		return err
	}
	if err := resource.Set("subscriptions_deleted", cast.ToInt(output.Attributes["SubscriptionsDeleted"])); err != nil {
		return err
	}
	// ... More attributes are set here ...

	return nil
}
```
