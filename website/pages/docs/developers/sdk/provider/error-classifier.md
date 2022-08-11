import { Callout } from 'nextra-theme-docs'

# Error Classifier

**Error Classifier** allows the provider to classify errors it produces during table execution, and return them as diagnostics to the user.

It's an optional feature to be configured in the provider:

```go
	// Classifier function may return empty slice if it cannot meaningfully convert the error into diagnostics. In this case
	// the error will be converted by the SDK into diagnostic at ERROR level and RESOLVING type.
	ErrorClassifier func(meta schema.ClientMeta, resource string, err error) []diag.Diagnostic
```

This way a specific error from a [Fetch Resolver](../table/fetch-resolvers) can be classified with a severity level and type, to be reported differently to the user.

<Callout type="info">

The error classifier is not able to ignore errors. Ignoring errors should be done at a [Table](../table/overview) level, using an `IgnoreErrorFunc`.

</Callout>
