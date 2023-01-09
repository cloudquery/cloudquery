---
title: How to Automatically Update Plugins Using Renovate
tag: tutorial
description: >-
  How to setup Renovate to update plugins
author: cloudmatt # Only renders author block for people in content/team.ts
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"
import { Callout } from 'nextra-theme-docs'

<HowToGuideHeader/>

In this guide we will discuss how to leverage [Renovate](https://github.com/renovatebot/renovate) to keep all your source and destination plugins on the latest version.

## Introduction to Renovate

[Renovate](https://github.com/renovatebot/renovate) is a tool for automated dependency management for your code repositories to help automate away the tasks of manual dependency versioning. For more information on the features and why you should use Renovate, please read [this](https://github.com/renovatebot/renovate#why-use-renovate).

This guide covers how to leverage Renovate to keep your CloudQuery source and destination plugins up-to-date but is not meant to be a comprehensive tutorial of Renovate. For instructions on how to setup Renovate, please reference their [getting started](https://docs.renovatebot.com/getting-started/running/) documentation. This how-to guide assumes you have Renovate configured and running already on your git based repository.

## Walkthrough

### Step 1: Identify Regex Capture Groups

While Renovate provides tons of built-in functionality to make most dependency management simple, it also provides a powerful functionality called 'RegexManager' which allows you to leverage custom regular expression statements to manually extract dependencies, their version, and updating mechanism. Click [here](https://docs.renovatebot.com/modules/manager/regex/) for more details on the specific configurations of RegexManager.

For our purposes, we're concerned about the following items:
- `datasourceTemplate`: datasource is where the package's versions are managed.
- `packageNameTemplate`:  this is the lookup name that is passed to the datasource.
- `depNameTemplate`: this is the 'friendly' name that is shown in renovate PRs.
- `extractVersionTemplate`: regular expression for discovering plugin version in the datasource.
- `versioningTemplate`: what versioning 'format' the datasource follows
- `fileMatch`: path to the `cloudquery.yaml` file in your repository.
  `matchStrings`: regular expression to find the plugins in your cloudquery.yaml file.

Using a simplistic configuration file that looks like:

```yaml copy
---
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "v5.4.0"
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v9.1.0"
  destinations: ["postgresql"]
  tables: ["*"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v2.0.1"
  write_mode: overwrite-delete-stale
``` 

You can see that we have three 'keys', meaning fields that are static across all plugins, we can leverage these to construct our custom regex statement for each plugin. These 'keys' will be used as regular expression named capture group. Those fields being:

- **kind**: what type of plugin this is (source or destination).
- **path**: the 'name' of the plugin (such as gcp or aws).
- **version**: the current version of the plugin.

### Step 2: Create Custom RegexManager

Now that we've identified the fields we need to extract from the configuration file to create the custom regex, next we need to construct the regex statement to extract those 'keys' into named capture groups.

For this it's recommended to use a regex testing tool to ensure your regex is actually working. [Regex101](https://regex101.com/) is a popular one but feel free to use any others as long as they support the PCRE2 regex engine. So taking our sample configuration file from above, we begin to play around with various regex until we meet the following criteria:

- a named capture group for the plugin type (source or destination) which we will reference as 'kind' in this tutorial.
- a named capture group for the plugin name (aws, gcp, etc..) which we will reference as 'plugin' in this tutorial.
- a named capture group for the plugin current version which **must** be named `currentValue`.

For our sample configuration file this regex will look like:

```shell copy
kind:\s(?<kind>.*)\nspec:\n\s{2}name:.*?\w\n\s{2}path:\scloudquery\/(?<plugin>.*)\n\s{2}version:\s\"?v(?<currentValue>.*\d)\"?\n
```
This regular expression does the following:
- captures the `kind: (source|destination)` into a named capture group called 'kind`
- skips the `spec:` and `name:` lines of the file
- captures the `path: cloudqery/(gcp, aws, etc...)` into a named capture group called `plugin` while escaping the `cloudquery/` prefix.
- skips to the next line and captures the `version: "vX.X.X"` into the **required** named capture group `currentValue` while optionally checking if the string is quoted or not.


On the Regex101 site we can verify the regex is working by viewing the explanation section on the left side of the screen:

![screenshot of regex101](/images/blog/manage-plugins-with-renovate/renovate-howto-regex101.png)

**Important Note**: This regex is _specific_ to the example shown in this how-to guide, if your configuration file is formatted different (such as lines in different orders or indented differently like in a helm `values.yaml` file) then you'll need to adjust the regex accordingly.

### Step 3: Add Custom RegexManager to Renovate

The last thing we need to get this working is to simply add a block to the `regexManagers: []` in your `renovate.json` file in your code repository such as:

```json copy
{
  "regexManagers": [
    {
      "description": "Update Cloudquery Plugins",
      "fileMatch": ["^cloudquery.yaml$"],
      "matchStrings" [
        "kind:\\s(?<kind>.*)\\nspec:\\n\\s{2}name:.*?\\w\\n\\s{2}path:\\scloudquery\/(?<plugin>.*)\\n\\s{2}version:\\s\"?v(?<currentValue>.*\\d)\"?\\n"
      ],
      "datasourceTemplate": "github-releases",
      "packageNameTemplate": "cloudquery/cloudquery",
      "depNameTemplate": "{{kind}}-{{plugin}}",
      "extractVersionTemplate": "^plugins-{{kind}}-{{plugin}}-v(?<version>.*)$",
      "versioningTemplate": "semver"
    }
  ]
}
```

Things to be aware of when adding this regex to your `renovate.json` file:
- Any regular expression token (such as `\s` or `\n`) needs to be escaped in the configuration file as shown above. This does **not** apply to regular escapes such as escaping the `"` with a singular `\`.
- the `depNameTemplate` is the friendly name we'll see in PRs for the plugins. Our example simply uses the kind and plugin names from the capture group but feel free to use your own.
- the `extractVersionTemplate` must match the releases in the CloudQuery github repository which at the time of this guide being written follows this format: `plugins-source-aws-v9.1.0`. We simply use the named capture groups to substitute into the string, simply think of them as 'variables' in this context.
- the `packageNameTemplate` is simply the name of the GitHub repository for CloudQuery, this is just what the `github-releases` datasource expects as input.

Once you've done this and renovate successfully runs, you should see the following the next time a plugin is released on github (with different labels based on your git workflow and CI):

![pull request tab in GitHub](/images/blog/manage-plugins-with-renovate/renovate-pr-page.png)

![example pull request](/images/blog/manage-plugins-with-renovate/renovate-pr-example.png)

![files changed](/images/blog/manage-plugins-with-renovate/renovate-pr-changedfiles.png)

## Summary

In this post we've discussed how you can leverage Renovate for automatic plugin updates to ensure you're always running the latest versions of your source and destination plugins without having to manually update.