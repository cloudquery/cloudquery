# Creating a New Plugin

The best way to learn how to create a new plugins is to look at the following examples:

- [Test Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/test) - Source plugin boilerplate code with one table.
- [Test Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/test) - Destination plugin boilerplate code.

More real world examples are:
- [GCP Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp)
- [PostgreSQL Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/postgresql)

All other official source and destination plugins can be found [here](https://github.com/cloudquery/cloudquery/tree/main/plugins)

## Naming Conventions

Community plugins use the following Github repository naming conventions:

 - `org/cq-source-<name>` for source plugins
 - `org/cq-destination-<name>` for destination plugins

A community plugin using this convention can be imported in a config by using: 

```yaml
kind: source
spec:
  path: org/name
``` 

for source plugins, or

```yaml
kind: destination
spec:
  path: org/name
```

for destination plugins. 

Names generally contain no dashes or underscores. So for example, if you are developing a source plugin for a new cloud service called Cloud Widgets, you should create the plugin repository under `org/cq-source-cloudwidgets`.

Official plugins, in contrast, are contained in the `cloudquery/cloudquery` monorepo. By convention, they can be imported using a special path `cloudquery/<name>`, e.g.:

```
kind: source
spec:
  path: cloudquery/aws
```