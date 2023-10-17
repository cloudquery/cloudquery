import AlicloudConfiguration from "../pages/docs/plugins/sources/alicloud/_configuration.md";
import AlicloudAuthentication from "../pages/docs/plugins/sources/alicloud/_authentication.md";
import AWSConfiguration from "../pages/docs/plugins/sources/aws/_configuration.md";
import AWSAuthentication from "../pages/docs/plugins/sources/aws/_authentication.md";
import AWSPricingConfiguration from "../pages/docs/plugins/sources/awspricing/_configuration.md";
import AzureAuthentication from "../pages/docs/plugins/sources/azure/_authentication.md";
import AzureConfiguration from "../pages/docs/plugins/sources/azure/_configuration.md";
import AzureDevopsAuthentication from "../pages/docs/plugins/sources/azuredevops/_authentication.md";
import AzureDevopsConfiguration from "../pages/docs/plugins/sources/azuredevops/_configuration.md";
import BitbucketAuthentication from "../pages/docs/plugins/sources/bitbucket/_authentication.md";
import BitbucketConfiguration from "../pages/docs/plugins/sources/bitbucket/_configuration.md";
import CloudflareAuthentication from "../pages/docs/plugins/sources/cloudflare/_authentication.md";
import CloudflareConfiguration from "../pages/docs/plugins/sources/cloudflare/_configuration.md";
import AZBlobAuthentication from "../pages/docs/plugins/destinations/azblob/_authentication.md";
import AZBlobConfiguration from "../pages/docs/plugins/destinations/azblob/_configuration.md";
import DuckDBConfiguration from "../pages/docs/plugins/destinations/duckdb/_configuration.md";
import GremlinConfiguration from "../pages/docs/plugins/destinations/gremlin/_configuration.md";
import PostgresDestConfiguration from "../pages/docs/plugins/destinations/postgresql/_configuration.md";
import SnowflakeConfiguration from "../pages/docs/plugins/destinations/snowflake/_configuration.md";
import SQLiteConfiguration from "../pages/docs/plugins/destinations/sqlite/_configuration.md";

export const components = {
  "sources-alicloud-configuration": <AlicloudConfiguration />,
  "sources-alicloud-authentication": <AlicloudAuthentication />,
  "sources-aws-configuration": <AWSConfiguration />,
  "sources-aws-authentication": <AWSAuthentication />,
  "sources-awspricing-configuration": <AWSPricingConfiguration />,
  "sources-azure-authentication": <AzureAuthentication />,
  "sources-azure-configuration": <AzureConfiguration />,
  "sources-azuredevops-authentication": <AzureDevopsAuthentication />,
  "sources-azuredevops-configuration": <AzureDevopsConfiguration />,
  "sources-bitbucket-authentication": <BitbucketAuthentication />,
  "sources-bitbucket-configuration": <BitbucketConfiguration />,
  "sources-cloudflare-authentication": <CloudflareAuthentication />,
  "sources-cloudflare-configuration": <CloudflareConfiguration />,
  "destinations-azblob-authentication": <AZBlobAuthentication />,
  "destinations-azblob-configuration": <AZBlobConfiguration />,
  "destinations-duckdb-configuration": <DuckDBConfiguration />,
  "destinations-gremlin-configuration": <GremlinConfiguration />,
  "destinations-postgresql-configuration": <PostgresDestConfiguration />,
  "destinations-snowflake-configuration": <SnowflakeConfiguration />,
  "destinations-sqlite-configuration": <SQLiteConfiguration />,
};
