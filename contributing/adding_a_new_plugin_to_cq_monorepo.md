# Adding New Plugin To CQ Monorepo

This guide relates only when you add a new plugin to the CloudQuery Monorepo.

There are number of steps to follow in order to add all the relevant CI and release processes to your plugin.

1. Add it to the [release please configuration file](../release-please-config.json).
2. Create a workflow file for it. See example [destination](../.github/workflows/dest_test.yml) and [source](../.github/workflows/source_gcp.yml) plugins.
3. Add the workflow file job name to the [wait-for-required-workflows workflow](../.github/workflows/wait_for_required_workflows.yml), so it will be enforced. For [example](https://github.com/cloudquery/cloudquery/blob/5c6e5a8eb5b8c6868336967dfe1a375cef5a792f/.github/workflows/wait_for_required_workflows.yml#L51).
4. Ensure there’s a `Version` var under `resources/plugin/plugin.go` so the version will embedded correctly by Go Releaser, example [here](https://github.com/cloudquery/cloudquery/blob/fb690589a1d2b7ed30f90744d156a6e5b0e57d66/plugins/destination/test/resources/plugin/plugin.go#L5). See also the relevant [Go Releaser configuration file](https://github.com/cloudquery/cloudquery/blob/812241697c644bdb1ae202bbadcb3baae456f788/plugins/.goreleaser.yaml#L12).
5. Add a `.goreleaser.yaml` file - [see example](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/.goreleaser.yaml).
6. If relevant, add it to the Website under [sources](../website/pages/docs/plugins/sources/overview.mdx) or [destinations](../website/pages/docs/plugins/destinations/overview.mdx).
7. If relevant, add it to the Website [sources _meta.json](../website/pages/docs/plugins/sources/_meta.json) or [destinations _meta.json](../website/pages/docs/plugins/destinations/_meta.json) file so that it displays correctly in the sidebar.
8. If relevant, add an entry in our [PR labeler](../.github/pr_labeler.yml).
9. Create a Sentry project for it under https://sentry.io/organizations/cloudquery-v2/projects/ and embed the correct DSN, for example see [here](https://github.com/cloudquery/cloudquery/blob/0e4b8dc53358388f8a1e61cad8ae8a1ab2f52342/plugins/source/azure/main.go#L8).
10. **Only relevant for big source plugins - after the initial PR is merged**, allow the plugin workflow file to access large runners via https://github.com/organizations/cloudquery/settings/actions/runner-groups/6.
11. **After the initial version of the plugin is released** add a “filler` entry for it in [here](https://github.com/cloudquery/cloudquery/blob/fb690589a1d2b7ed30f90744d156a6e5b0e57d66/.release-please-manifest.json#L29). This ensures we don’t get conflicts when creating multiple release PRs. More about this in this [issue](https://github.com/googleapis/release-please/issues/1502).
