# Adding New Plugin To CQ Monorepo

This small guide relates only when you add a new plugin to the CQ Monorepo.

There are number of steps to follow in order to add all the relevant CI and release processes to your plugin.

1) Add your plugin to the `plugins/destination` or `plugins/source` folder.
2) Add an entry to `.github/pr_labeler.yml`
3) Add an entry to `./github/workflow/dest_your_plugin.yml` for running tests and linting.
4) Add an entry to `.github/workflows/wait_for_required_workflows.yml`
5) Add an entry to `release-please-config.json`
