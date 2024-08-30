import { AuthType, DestinationConfig, Link } from '@cloudquery/plugin-config-ui-lib';

import { AdvancedConnectionFields } from './form/sections/advancedConnectionFields';
import { FormConnectionFields } from './form/sections/connectionFields';
import { migrateModeValues, writeModeValues } from './utils/formSchema';
import { pluginUiMessageHandler } from './utils/messageHandler';

const config: DestinationConfig = {
  name: 'mysql',
  type: 'destination',
  label: 'MySQL',
  docsLink: 'https://hub.cloudquery.io/plugins/destination/cloudquery/mysql/latest/docs',
  iconLink: 'images/logo.webp',
  steps: [
    {
      sections: [
        FormConnectionFields,
        AdvancedConnectionFields,
        {
          component: 'section',
          title: 'Sync Options',
          children: [
            {
              component: 'control-select-field',
              name: 'migrateMode',
              options: [...migrateModeValues],
              helperText: (
                <>
                  Specifies the migration mode to use when source tables are changed.{' '}
                  <Link
                    href="https://docs.cloudquery.io/docs/reference/destination-spec#migrate_mode"
                    pluginUiMessageHandler={pluginUiMessageHandler}
                  >
                    Read more
                  </Link>
                </>
              ),
              label: 'Migrate mode',
            },
            {
              component: 'control-select-field',
              name: 'writeMode',
              options: [...writeModeValues],
              helperText: (
                <>
                  Specifies the update method to use when inserting rows.{' '}
                  <Link
                    href="https://docs.cloudquery.io/docs/reference/destination-spec#write_mode"
                    pluginUiMessageHandler={pluginUiMessageHandler}
                  >
                    Read more
                  </Link>
                </>
              ),
              label: 'Write mode',
            },
          ],
        },
        {
          component: 'collapsible-section',
          title: 'Advanced Sync Options',
          children: [
            {
              component: 'control-number-field',
              name: 'batch_size',
              helperText:
                'Maximum number of items that may be grouped together to be written in a single write. Default is 10,000.',
              label: 'Batch size',
            },
            {
              component: 'control-number-field',
              name: 'batch_size_bytes',
              helperText:
                'Maximum size of items that may be grouped together to be written in a single write. Default is 100,000,000 = 100MB.',
              label: 'Batch size (bytes)',
            },
          ],
        },
      ],
      title: 'Configuration',
    },
  ],
  auth: [AuthType.OTHER],
  guide: {
    title: 'MySQL configuration',
    sections: [
      {
        bodies: [
          {
            text: (
              <>
                The MySQL destination lets you sync data from any CloudQuery source to a
                MySQL&#8209;compatible database.
              </>
            ),
          },
        ],
      },
      {
        header: 'Setup guide',
        bodies: [
          {
            text: `To allow CloudQuery network access to your MySQL instance, make sure the following
            CloudQuery IPs are in your firewall allowlist:`,
          },
          { code: '35.231.218.115' },
          { code: '35.231.72.234' },
        ],
      },
    ],
  },
  errorCodes: {
    INVALID_DSN:
      'The connection string (DSN) is invalid or in an incorrect format. Please check and correct your connection details.',
    CONNECT_FAILED:
      'Failed to establish a connection to the MySQL database. This is rare and might indicate a driver issue.',
    DEFAULT_DATABASE_FAILED:
      "Unable to determine the default database. Please ensure you've specified a database name in your connection string.",
    QUERY_VERSION_FAILED:
      'Failed to retrieve the MySQL version. This might indicate restricted permissions or a connection issue.',
    UNREACHABLE: 'The MySQL server is unreachable. Check your host, port, and network settings.',
    ACCESS_DENIED: 'Access denied. The provided username or password is incorrect.',
    UNKNOWN_DATABASE:
      "The specified database does not exist. Please check your database name and ensure it's created on the server.",
    PING_FAILED:
      'Failed to ping the MySQL server. This might indicate network issues or server unavailability.',
    LIST_FAILED: 'Failed to list databases. This might be due to insufficient permissions.',
  },
};

export default config;
