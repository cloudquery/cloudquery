# Azure Source Plugin Contribution Guide

## Overview

This plugin uses code-generation heavily to speed up development and take advantage of the [Azure SDK](https://github.com/Azure/azure-sdk-for-go) which is in itself auto-generate (for the most part).



## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Azure API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
