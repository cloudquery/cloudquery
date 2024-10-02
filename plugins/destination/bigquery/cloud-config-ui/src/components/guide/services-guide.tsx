import { Link, RenderGuide } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

export function ServicesGuide() {
  return (
    <RenderGuide
      pluginUiMessageHandler={pluginUiMessageHandler}
      sections={[
        {
          bodies: [
            {
              text: 'Select the services that you want to sync the data for. This will affect the API endpoints that will be queried and the tables created in your destination database.',
            },
            {
              text: 'The most popular services are listed on top. Note that some services may take a while to sync depending on the amount of resources used.',
            },
            {
              text: (
                <>
                  See the{' '}
                  <Link
                    href="https://hub.cloudquery.io/plugins/source/cloudquery/gcp/latest/docs"
                    pluginUiMessageHandler={pluginUiMessageHandler}
                  >
                    GCP Plugin Documentation
                  </Link>{' '}
                  to see the full list of tables and their schema.
                </>
              ),
            },
          ],
        },
      ]}
    />
  );
}
