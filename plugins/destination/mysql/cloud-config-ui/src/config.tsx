import { AuthType, DestinationConfig, Link } from '@cloudquery/plugin-config-ui-lib';

import { migrateModeValues, tlsModeValues, writeModeValues } from './utils/formSchema';
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
        {
          component: 'section',
          title: 'Connect to your database',
          subtitle: 'Set up a connection to your MySQL instance.',
          children: [
            {
              component: 'control-exclusive-toggle',
              name: '_connectionType',
              options: [
                {
                  label: 'Regular setup',
                  value: 'fields',
                },
                {
                  label: 'Connection string',
                  value: 'string',
                },
              ],
            },
            {
              component: 'control-secret-field',
              name: 'connection_string',
              helperText:
                'Connection string to connect to the database. E.g. user:password@localhost:3306/dbname?tls=preferred\u0026readTimeout=1s\u0026writeTimeout=1s',
              label: 'Connection string',
              shouldRender: (values: any) => values._connectionType === 'string',
            },
            {
              component: 'control-text-field',
              name: 'host',
              helperText:
                'Host to connect to. E.g. 1.2.3.4 or mydb.host.com. Optional, defaults to empty.',
              label: 'Host',
              shouldRender: (values: any) => values._connectionType === 'fields',
            },
            {
              component: 'control-number-field',
              name: 'port',
              helperText: 'Port to connect to. Optional, defaults to empty.',
              label: 'Port',
              shouldRender: (values: any) => values._connectionType === 'fields',
            },
            {
              component: 'control-text-field',
              name: 'database',
              helperText:
                'Name of the MySQL database you want to connect to. Optional, defaults to empty.',
              label: 'Database',
              shouldRender: (values: any) => values._connectionType === 'fields',
            },
            {
              component: 'control-text-field',
              name: 'username',
              helperText: 'Username to use when authenticating. Optional, defaults to empty.',
              label: 'Username',
              shouldRender: (values: any) => values._connectionType === 'fields',
            },
            {
              component: 'control-secret-field',
              name: 'password',
              label: 'Password',
              helperText: 'Password to use when authenticating. Optional, defaults to empty.',
              shouldRender: (values: any) => values._connectionType === 'fields',
            },
          ],
        },
        {
          component: 'collapsible-section',
          title: 'Advanced Connection Options',
          defaultExpanded: false,
          shouldRender: (values: any) => values._connectionType === 'fields',
          children: [
            {
              component: 'control-boolean-field',
              name: 'tcp',
              label: 'TCP',
              type: 'toggle',
              helperText:
                'If true, will enable connection over TCP to the server. Optional, defaults to true.',
            },
            {
              component: 'control-boolean-field',
              name: 'connectionParams.tls',
              label: 'TLS',
              type: 'toggle',
              helperText:
                'If true, will enable TLS/SSL encrypted connection to the server. Optional, defaults to false.',
            },
            {
              component: 'control-select-field',
              name: 'connectionParams.tlsMode',
              helperText:
                'SSL connections to encrypt client/server communications using TLS protocols for increased security.',
              label: 'TLS Mode',
              shouldRender: (values: any) => !!values.connectionParams.tls,
              options: [...tlsModeValues],
            },
            {
              component: 'control-boolean-field',
              name: 'connectionParams.parseTime',
              label: 'Parse Time',
              type: 'toggle',
              helperText:
                'If true, changes the output type of DATE and DATETIME values to time.Time instead of []byte / string. Optional, defaults to false.',
            },
            {
              component: 'control-text-field',
              name: 'connectionParams.loc',
              helperText:
                'Sets the location for time.Time values. "Local" sets the system\'s location. Optional, defaults to UTC.',
              label: 'Location',
              shouldRender: (values: any) => !!values.connectionParams.parseTime,
            },
            {
              component: 'control-text-field',
              name: 'connectionParams.charset',
              helperText:
                'Sets the charset used for client-server interaction. Multiple charsets can be configured with comma separation (ex. utf8mb4,utf8). Optional, defaults to utf8mb4.',
              label: 'Charset',
            },
            {
              component: 'control-number-field',
              name: 'connectionParams.timeout',
              helperText:
                'Timeout for establishing connections, aka dial timeout. Value is in seconds. Optional, defaults to 0.',
              label: 'Timeout',
            },
            {
              component: 'control-number-field',
              name: 'connectionParams.readTimeout',
              helperText: 'I/O read timeout. Value is in seconds. Optional, defaults to 0.',
              label: 'Read Timeout',
            },
            {
              component: 'control-number-field',
              name: 'connectionParams.writeTimeout',
              helperText: 'I/O write timeout. Value is in seconds. Optional, defaults to 0.',
              label: 'Write Timeout',
            },
          ],
        },
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
          defaultExpanded: false,
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
