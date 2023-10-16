import AlicloudConfiguration from "../pages/docs/plugins/sources/alicloud/_configuration.md";
import AlicloudAuthentication from "../pages/docs/plugins/sources/alicloud/_authentication.md";
import AWSConfiguration from "../pages/docs/plugins/sources/aws/_configuration.md";
import AWSAuthentication from "../pages/docs/plugins/sources/aws/_authentication.md";

import FirehoseConfiguration from "../pages/docs/plugins/destinations/firehose/_configuration.md";
import FirehoseAuthentication from "../pages/docs/plugins/destinations/firehose/_authentication.md";

export const components = {
  "sources-alicloud-configuration": <AlicloudConfiguration />,
  "sources-alicloud-authentication": <AlicloudAuthentication />,
  "sources-aws-configuration": <AWSConfiguration />,
  "sources-aws-authentication": <AWSAuthentication />,

  "destinations-firehose-configuration": <FirehoseConfiguration />,
  "destinations-firehose-authentication": <FirehoseAuthentication />,
};
