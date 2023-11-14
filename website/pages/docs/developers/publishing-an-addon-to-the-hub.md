---
title: Publishing an Addon to the Hub
description: Learn how to publish an addon to the Hub and have it listed in https://hub.cloudquery.io/
---

# Publishing an Addon to the Hub

With the announcement of [CloudQuery Hub](/blog/announcing-cloudquery-new-hub), we are excited to see the community contribute addons, such as transformations and dashboards, to the Hub. This guide will walk you through the process of publishing an addon to the Hub.

## Prerequisites

- You have created a [CloudQuery Cloud](https://cloud.cloudquery.io/) account and completed the onboarding process to create a team
- You have created the addon you'd like to publish on https://cloud.cloudquery.io/ under the relevant team
- You have the [CloudQuery CLI](/docs/quickstart) installed (version >= `v3.27.1`)
- You are authenticated to [CloudQuery Cloud](https://cloud.cloudquery.io/) using the `cloudquery login` command or an API key

## Publishing an Addon

1. (Optional, recommended) In the root directory of your addon repository run `git tag v1.0.0` to tag the version you're about to publish (replace `v1.0.0` with the version you'd like to publish)
2. (Optional, recommended) Run `git push origin v1.0.0` to push the tag
3. Create a `manifest.json` file that describes the addon, the path to a zip containing its files and the path to its documentation and changelog in markdown format. Here is an example:
   ```json copy filename="manifest.json"
   {
     "schema_version": 1,
     "type": "addon",
     "team_name": "my_team",
     "addon_name": "example",
     "addon_type": "visualization",
     "addon_format": "zip",
     "message": "./changelog.md",
     "doc": "./readme.md",
     "path": "./test.zip",
     "plugin_deps": ["cloudquery/source/test@v1.0.0"],
     "addon_deps": []
   }
   ```
    
   The `plugin_deps` field describes the plugins that this addon depends on, if any. The format is `<team>/<kind>/<name>@<version>`.
   Similarly, the `addon_deps` field describes the addons that this addon depends on, if any. The format is `<team>/<type>/<name>@<version>`.

4. Run `cloudquery addon publish /path/to/manifest.json v1.0.0` to publish a draft version of the addon (replacing `v1.0.0` with the version you want to publish). The version will show up under the versions tab of your addon in <https://cloud.cloudquery.io>. As long as the version is in draft it's mutable and you can re-package the addon and publish it again.
5. Once you're ready, run `cloudquery addon publish /path/to/manifest.json v1.0.0 --finalize` to publish a non-draft version of the addon. This version will be immutable and will show up in <https://hub.cloudquery.io/>. Allow up to 1 hour for the Hub to reflect the changes, and please allow time for the CloudQuery team to review the addon before it's published.
