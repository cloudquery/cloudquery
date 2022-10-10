# Generating resources

Adding resources to a plugins can sometimes be a tedious task, some resources can have more than hundreds of fields and relations, and adding them all can
take a long time. To remedy this issue, we provide code generation utilities as part of our [plugin-sdk](https://github.com/cloudquery/plugin-sdk). Code generation allows to easily generate more of the boilerplate code for tables from Go code.

## Examples

The best example would be to checkout how this is done in our of our official plugins such as [gcp](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/codegen)
