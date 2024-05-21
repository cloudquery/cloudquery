# Creating a New Plugin

There are two main types of CloudQuery plugins:
 - **source** plugins read data from a cloud provider, and
 - **destination** plugins define a destination to write source data into. This is usually a database or file system.

## Writing a Source Plugin

See the following links for more information on writing source plugins (we will add more resources here over time):

- [Creating a New Plugin Docs](https://www.cloudquery.io/docs/developers/creating-new-plugin)
- [Create a Source Plugin Video Tutorial](https://www.youtube.com/watch?v=3Ka_Ob8E6P8)

## Creating your first Premium Source Plugin as part of onboarding

(this section is based on a [real PR review](https://github.com/cloudquery/cloudquery-private/pull/2517) of the first plugin of a staff onboarding)

- The scaffolding tool recommended in [the guide](https://docs.cloudquery.io/docs/developers/creating-new-plugin/go-source#incremental-tables) is meant for public plugins, so it gets some things wrong, but it's still an efficient starting point. Follow the next bullets to correct things.
- Follow [these](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/contributing/adding_a_new_plugin_to_cq_monorepo.md ) steps, which are not in the guide. 
- The generated `Makefile` lacks some targets. Use some [existing source plugin's Makefile](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/wiz/Makefile#L1) for inspiration.
- Scaffolding of `README.md` doesn't capitalise the plugin name, and it builds a large `README.md`. It also adds badges that aren't necessary. Separate usage instructions into an `docs/overview.md` (like [here](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/wiz/docs/overview.md)). This latter file is what gets published to our Hub and should explain to users how to use the plugin.
- Don't implement retryability and backoff of requests. Use either https://github.com/hashicorp/go-retryablehttp or https://github.com/avast/retry-go. Grep the codebase for examples and follow the consensus.
- Make all Table's [paid](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/github/resources/plugin/tables.go#L57).
- Add usage reporting. See: [this](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/github/resources/plugin/client.go#L109), [this](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/github/resources/plugin/client.go#L72) and [this](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/github/resources/plugin/client.go#L85).
- If the table you're implementing has a PK, then [add a transformation](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/azure/resources/services/advisor/suppressions.go#L19) to your Table's transformer that sets the field as PK.
- Add a [Sentry DSN](https://github.com/cloudquery/cloudquery-private/blob/4f2ce0c9c5911828f997e10a9c9023fec4574e63/plugins/source/aws/main.go#L11) to your plugin. If you don't have Sentry access you can either request access or request a DSN to be created for the new plugin.
- There's a team convention to make table names plural.
- When possible, authentication parameters for the plugin should be fed via the spec (e.g. not via env variables or a file), because it would be cumbersome/impossible to supply this on cloud syncs. Sometimes this is impractical, e.g. AWS uses `~/.aws`.
- Err on the side of simplicity: don't make an initial version of a plugin more configurable than it needs to be for an MVP; let users consume it and add issues.
- When coding the internal API client, make the HTTP requests context-aware; the plugin client functions have a `ctx` argument.
- Scaffolding of `.gitignore` duplicates entries from the top-level one; all that is needed is the binary name.
- Spec parameters should be `snake-case`.

## Writing a Destination Plugin

This section is a work in progress. For now, raise an issue or reach out to us on Discord if 
you would like to request (or work on) a new destination plugin.  
