# Adding a New Plugin To CQ Monorepo

This guide relates only when you add a new plugin to the CloudQuery Monorepo.

There are number of steps to follow in order to add all the relevant CI and release processes to your plugin.

1. Add it to the [release please configuration file](../release-please-config.json).
2. Create a workflow file for it. See example [destination](../.github/workflows/dest_test.yml) and [source](../.github/workflows/source_xkcd.yml) plugins.
3. If relevant, add an entry in our [PR labeler](../.github/pr_labeler.yml).
4. **After the initial version of the plugin is released** add a "filler" entry for it in [here](https://github.com/cloudquery/cloudquery/blob/fb690589a1d2b7ed30f90744d156a6e5b0e57d66/.release-please-manifest.json#L29). This ensures we don’t get conflicts when creating multiple release PRs. More about this in this [issue](https://github.com/googleapis/release-please/issues/1502).
