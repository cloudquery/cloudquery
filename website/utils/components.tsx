import AirtableAuthentication from "../pages/docs/plugins/sources/airtable/_authentication.md";
import AirtableConfiguration from "../pages/docs/plugins/sources/airtable/_configuration.md";
import BitbucketAuthentication from "../pages/docs/plugins/sources/bitbucket/_authentication.md";
import BitbucketConfiguration from "../pages/docs/plugins/sources/bitbucket/_configuration.md";
import SquareAuthentication from "../pages/docs/plugins/sources/square/_authentication.md";
import SquareConfiguration from "../pages/docs/plugins/sources/square/_configuration.md";
import TypeformAuthentication from "../pages/docs/plugins/sources/typeform/_authentication.md";
import TypeformConfiguration from "../pages/docs/plugins/sources/typeform/_configuration.md";


export const components = {
  "sources-airtable-authentication": <AirtableAuthentication />,
  "sources-airtable-configuration": <AirtableConfiguration />,
  "sources-bitbucket-authentication": <BitbucketAuthentication />,
  "sources-bitbucket-configuration": <BitbucketConfiguration />,
  "sources-square-authentication": <SquareAuthentication />,
  "sources-square-configuration": <SquareConfiguration />,
  "sources-typeform-authentication": <TypeformAuthentication />,
  "sources-typeform-configuration": <TypeformConfiguration />,
};
