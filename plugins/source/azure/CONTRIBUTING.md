# Azure Source Plugin Contribution Guide

## Overview

This plugin uses code-generation heavily to speed up development and take advantage of the [Azure SDK](https://github.com/Azure/azure-sdk-for-go) which is in itself auto-generate (for the most part).

There are two steps of auto-generation:

### Discovery and auto-generating recipes

The first problem with generating all azure tables is even discovering which packages are available and then generating the recipes under `codegen2/`. For this we use a multi-stage discovery and generation process which for the most part should run periodically to update code but developer would rarely change that.

Process:

1) `scripts/update_azure_subpackages.sh` will curl `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk` and update `go.mod` with all relevant sub-packages and download those modules locally.
2) `codegen0` will then parse the source code using go `ast` to find all root `NewXClients` which provides `NewListPager` function and will generate all recipes for `codegen1/recipes/*`.
3) This intermediate step is used because not everything is easily available when parsing the AST so we need to be able to reflect some of the types and then generate the final recipes to be located under `codegen2/recipes/*`

### Generating the actual tables

To generate the actual tables take a look at the recipes under `codegen2/recipes/*` and add your own based on other recipes.


## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Azure API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
