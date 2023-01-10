# Creating a New Plugin

The best way to learn how to create a new plugins is to look at the following examples:

- [Test Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/test) - Source plugin boilerplate code with one table.
- [Test Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/test) - Destination plugin boilerplate code.

More real world examples are:

- [GCP Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp)
- [PostgreSQL Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/postgresql)

Other source and destination plugins to reference can be found [here](https://github.com/cloudquery/cloudquery/tree/main/plugins)

## Naming Conventions

Community plugins use the following GitHub repository naming conventions:

- `org/cq-source-<name>` for source plugins
- `org/cq-destination-<name>` for destination plugins

A community plugin using this convention can be imported in a config by using:

```yaml copy
kind: source
spec:
  path: org/name
```

for source plugins, or

```yaml copy
kind: destination
spec:
  path: org/name
```

for destination plugins.

Names should not contain dashes or underscores. So for example, if you are developing a source plugin for a new cloud service called Cloud Widgets, you should create the plugin repository under `org/cq-source-cloudwidgets`.

Official plugins, in contrast, are contained in the [CloudQuery repository](https://github.com/cloudquery/cloudquery/tree/main/plugins). By convention, they can be imported using a special path `cloudquery/<name>`, e.g.:

```yaml copy
kind: source
spec:
  path: cloudquery/aws
```
