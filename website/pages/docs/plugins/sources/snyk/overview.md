# Snyk Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `snyk`)}/>

The CloudQuery Snyk plugin pulls configuration out of Snyk
resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Snyk, `cloudquery` needs to be authenticated.
An API key is required for authentication.
See [Authentication for API](https://docs.snyk.io/snyk-api-info/authentication-for-api) for more information.
