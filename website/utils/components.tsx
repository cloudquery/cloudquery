import AlicloudConfiguration from "../pages/docs/plugins/sources/alicloud/_configuration.md";
import AlicloudAuthentication from "../pages/docs/plugins/sources/alicloud/_authentication.md";
import AWSConfiguration from "../pages/docs/plugins/sources/aws/_configuration.md";
import AWSAuthentication from "../pages/docs/plugins/sources/aws/_authentication.md";

import AZBlobAuthentication from "../pages/docs/plugins/destinations/azblob/_authentication.md";
import AZBlobConfiguration from "../pages/docs/plugins/destinations/azblob/_configuration.md";

export const components = {
  "sources-alicloud-configuration": <AlicloudConfiguration />,
  "sources-alicloud-authentication": <AlicloudAuthentication />,
  "sources-aws-configuration": <AWSConfiguration />,
  "sources-aws-authentication": <AWSAuthentication />,

  "destinations-azblob-authentication": <AZBlobAuthentication />,
  "destinations-azblob-configuration": <AZBlobConfiguration />,
};
