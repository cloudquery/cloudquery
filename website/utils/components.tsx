import AlicloudConfiguration from "../pages/docs/plugins/sources/alicloud/_configuration.md";
import AlicloudAuthentication from "../pages/docs/plugins/sources/alicloud/_authentication.md";
import AWSConfiguration from "../pages/docs/plugins/sources/aws/_configuration.md";
import AWSAuthentication from "../pages/docs/plugins/sources/aws/_authentication.md";
import PostgresDestConfiguration from "../pages/docs/plugins/destinations/postgresql/_configuration.md";
import GremlinConfiguration from "../pages/docs/plugins/destinations/gremlin/_configuration.md";
import SnowflakeConfiguration from "../pages/docs/plugins/destinations/snowflake/_configuration.md";

export const components = {
  "sources-alicloud-configuration": <AlicloudConfiguration />,
  "sources-alicloud-authentication": <AlicloudAuthentication />,
  "sources-aws-configuration": <AWSConfiguration />,
  "sources-aws-authentication": <AWSAuthentication />,

  "destinations-postgresql-configuration": <PostgresDestConfiguration />,
  "destinations-gremlin-configuration": <GremlinConfiguration />,
  "destinations-snowflake-configuration": <SnowflakeConfiguration />,
};
