import { Link, RenderGuide } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

const getServicesSelection = (editMode: boolean) => [
  {
    header: 'Step 1: Select regions',
    bodies: [
      {
        text: 'By default, all enabled AWS regions will be synced. Selecting regions to the left will filter down the regions that will be synced.',
      },
    ],
  },
  {
    header: 'Step 2: Select services',
    bodies: [
      {
        text: 'Select the services that you want to sync the data for. This will affect the API endpoints that will be queried and the tables created in your destination database.',
      },
      {
        text: 'By default, all enabled AWS services will be synced. Selecting services to the left will filter down the tables that will be synced.',
      },
      {
        text: 'The most popular services are listed on top. Note that some services may take a while to sync depending on the amount of resources used.',
      },
      {
        text: (
          <>
            See the{' '}
            <Link
              href="https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs"
              pluginUiMessageHandler={pluginUiMessageHandler}
            >
              AWS Plugin Documentation
            </Link>{' '}
            to see the full list of tables and their schema.
          </>
        ),
      },
    ],
  },
  {
    header: 'Step 3: Test the connection and submit',
    bodies: [
      {
        text: `Click the Test Connection and ${editMode ? 'update' : 'create'} source button to check if CloudQuery can connect and to submit the configuration.`,
      },
    ],
  },
];

export function AWSSelectServices({ editMode }: { editMode: boolean }) {
  const SERVICES_SECTION = getServicesSelection(editMode);

  return (
    <RenderGuide pluginUiMessageHandler={pluginUiMessageHandler} sections={SERVICES_SECTION} />
  );
}
