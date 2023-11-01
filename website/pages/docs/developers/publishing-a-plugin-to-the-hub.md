---
title: Publishing a Plugin to the Hub
description: Learn how to publish a plugin to the Hub and have it listed in https://hub.cloudquery.io/
---

# Publishing a Plugin to the Hub

With the announcement of [CloudQuery Hub](blog/announcing-cloudquery-new-hub), we are excited to see the community contribute plugins to the Hub. This guide will walk you through the process of publishing a plugin to the Hub.

:::callout
Publishing to the CloudQuery Hub is only supported for Go plugins at the moment. We are working on adding support for other languages.
:::

## Prerequisites

- You have created a [CloudQuery Cloud](https://cloud.cloudquery.io/) account and completed the onboarding process to create a team
- You have created the plugin you'd like to publish on https://cloud.cloudquery.io/ under the relevant team
- You have the [CloudQuery CLI](https://cloudquery.io/docs/cli/installation) installed (version >= `v3.27.1`)
- The plugin you'd like to publish is written in Go, and uses an SDK version >= `v4.17.1`
- The plugin you'd like to publish is initialized using the plugin's name, team and kind. See example [here](https://github.com/cloudquery/cloudquery/blob/b7ef6f6ed8948272a429f35614fa28559397227a/plugins/source/test/resources/plugin/plugin.go#L15)
- You are authenticated to [CloudQuery Cloud](https://cloud.cloudquery.io/) using the `cloudquery login` command

## Publishing a Plugin

1. (Optional, recommended) In the root directory of your plugin repository run `git tag v1.0.0` to tag the version you're about to publish (replace `v1.0.0` with the version you'd like to publish)
2. (Optional, recommended) Run `git push origin v1.0.0` to push the tag
3. Run `go run main.go package --docs-dir docs -m 'feat: Initial release' v1.0.0 .` to package the plugin. `v1.0.0` should match the tag you created in step 1. The `-m` specifies the changelog message that will be used in the release notes and it supports markdown. See example [here](https://hub.cloudquery.io/plugins/source/cloudquery/alicloud/v4.0.14/versions). `docs` should be a directory containing markdown files that serve as documentation for the plugin. Read more about the documentation format [here](#documentation format).
4. Run `cloudquery plugin publish` to publish a draft version of the plugin. The version will show up under the versions tab of your plugin in <https://cloud.cloudquery.io>. As long as the version is in draft it's mutable and you can re-package the plugin and publish it again
5. Once you're ready run `cloudquery plugin publish -f` to publish a non draft version of the plugin. This version will be immutable and will show up in <https://hub.cloudquery.io/>. Allow up to 1 hour for the Hub to reflect the changes

## Documentation Format

- The only documentation format supported at the moment is markdown, and the `cloudquery publish` command will only upload markdown files with the `.md` extension
- You can have multiple markdown files as documentation. The files will be concatenated in alphabetical order, and if one of the files is named `overview.md` it will show up first
- The markdown filename will be title cased when display in the Hub. For example `overview.md` will be displayed as `Overview`
- HTML tags are not supported in the markdown files and will be ignored
- Relative assets are not support e.g. `./assets/logo.png`. We recommend using absolute URLs for assets e.g. `https://raw.githubusercontent.com/<owner>/<repo>/main/assets/logo.png` in case you have the assets on GitHub