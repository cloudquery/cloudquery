# Provider Configuration

The `provider.Provider` struct looks like:

```go
type Provider struct {
	// Name of plugin
	Name string
	// Version of the provider
	Version string
	// Configure the provider and return context
	Configure func(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, error)
	// ResourceMap is all resources supported by this plugin
	ResourceMap map[string]*schema.Table
	// Configuration decoded from configure request
	Config func() Config
	// Logger to call, this logger is passed to the serve.Serve Client, if not define Serve will create one instead.
	Logger hclog.Logger
	// ErrorClassifier allows the provider to classify errors it produces during table execution, and return them as diagnostics to the user.
	// Classifier function may return empty slice if it cannot meaningfully convert the error into diagnostics. In this case
	// the error will be converted by the SDK into diagnostic at ERROR level and RESOLVING type.
	ErrorClassifier func(meta schema.ClientMeta, resource string, err error) []diag.Diagnostic

	// ...internal fields...
```

(Full definition [in repository](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/provider.go))

When creating a provider instance, the SDK decodes the configuration (from the user) into the data structure returned from the `Config` method.

```go
// Config Every provider implements a resources field we only want to extract that in fetch execution
type Config interface {
	// Example returns a configuration example (with comments) so user clients can generate an example config
	Example() string
}
```

`Config` is an `interface` and it's supposed to be owned by the provider, containing necessary configuration values.
The interface requires an `Example() string` method to be implemented, to generate an example config to use in `cloudquery init`.

After configuration is decoded into the data structure provided by `Config`, the method set in `Configure` is called.
This method takes the decoded config and initializes a client for the provider.

The returned client is then used per [table](../table/overview) and if a [multiplexer](../table/multiplexer-and-deletefilter) is enabled it's passed on to the multiplexer to be copied.
