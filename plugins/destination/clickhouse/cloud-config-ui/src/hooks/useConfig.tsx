import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { Link, AuthType, DestinationConfig } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import {
  compressValues,
  connectionOpenStrategyValues,
  connectionTypeValues,
  migrateModeValues,
  writeModeValues,
} from '../utils/constants';
import { convertConnectionStringToFields } from '../utils/convertConnectionStringToFields';
import { pluginUiMessageHandler } from '../utils/messageHandler';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
}

export const useConfig = ({ initialValues }: Props): DestinationConfig => {
  const url = initialValues?.spec?.connection_string || '';
  const connectionObj: Record<string, any> = convertConnectionStringToFields(url);

  return useMemo(
    () => ({
      debug: true,
      name: 'clickhouse',
      type: 'destination',
      label: 'ClickHouse',
      docsLink: 'https://hub.cloudquery.io/plugins/destination/cloudquery/clickhouse/latest/docs',
      iconLink: 'images/logo.webp',
      steps: [
        {
          children: [
            {
              component: 'section',
              title: 'Connect to your database',
              subtitle: 'Set up a connection to your ClickHouse instance.',
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
                  schema: yup
                    .string()
                    .oneOf(connectionTypeValues)
                    .default(connectionTypeValues[1])
                    .required(),
                },
                {
                  component: 'control-secret-field',
                  name: 'connection_string',
                  helperText:
                    'Connection string to connect to the database. E.g. clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms',
                  label: 'Connection string',
                  shouldRender: (values: any) => values._connectionType === 'string',
                  schema: yup
                    .string()
                    .default(url)
                    .when('_connectionType', {
                      is: 'string',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-multi-select',
                  name: 'hosts',
                  helperText: 'Hosts to connect to. E.g. 1.2.3.4:9000 or mydb.host.com:9000.',
                  label: 'Hosts',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .array()
                    .of(
                      yup
                        .string()
                        .max(253)
                        .matches(
                          /^(https?:\/\/)?[^:]+:(\d+)$/,
                          'Must be a valid host with port number, optionally starting with http:// or https://.',
                        ),
                    )
                    .default(connectionObj.hosts ?? [])
                    .when('_connectionType', {
                      is: 'fields',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.min(1),
                    }),
                },
                {
                  component: 'control-secret-field',
                  name: 'database',
                  helperText: 'Name of the ClickHouse database you want to connect to.',
                  label: 'Database',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(63)
                    .default(connectionObj.database ?? '')
                    .when('_connectionType', {
                      is: 'fields',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-secret-field',
                  name: 'username',
                  helperText: 'Username to use when authenticating.',
                  label: 'Username',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(63)
                    .default(connectionObj.username ?? '')
                    .when('_connectionType', {
                      is: 'fields',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-secret-field',
                  name: 'password',
                  label: 'Password',
                  helperText: 'Password to use when authenticating.',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(63)
                    .default(connectionObj.password ?? '')
                    .when('_connectionType', {
                      is: 'fields',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
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
                  name: 'connectionParams.secure',
                  label: 'SSL/TLS',
                  type: 'toggle',
                  helperText:
                    'If true, will enable SSL/TLS encrypted connection to the server. Optional, defaults to false.',
                  schema: yup.bool().default(connectionObj.connectionParams?.secure ?? false),
                },
                {
                  component: 'control-boolean-field',
                  name: 'connectionParams.skip_verify',
                  label: 'Skip certificate verification',
                  type: 'toggle',
                  helperText:
                    'If true, will skip certificate verification. Optional, defaults to false.',
                  schema: yup.bool().default(connectionObj.connectionParams.skip_verify ?? false),
                },
                {
                  component: 'control-number-field',
                  name: 'connectionParams.dial_timeout',
                  helperText:
                    'Timeout for establishing connections, aka dial timeout. Value is in milliseconds. Optional, defaults to 3000ms (or 3 seconds).',
                  label: 'Timeout',
                  schema: yup
                    .number()
                    .integer()
                    .default(connectionObj.connectionParams?.dial_timeout ?? 0),
                },
                {
                  component: 'control-select-field',
                  name: 'connectionParams.connection_open_strategy',
                  helperText:
                    'Strategy for choosing the order in which to connect to the servers. Optional, defaults to `in_order`',
                  label: 'Connection Open Strategy',
                  options: [...connectionOpenStrategyValues],
                  schema: yup
                    .string()
                    .oneOf(connectionOpenStrategyValues)
                    .default(
                      connectionObj.connectionParams?.connection_open_strategy ?? 'in_order',
                    ),
                },
                {
                  component: 'control-boolean-field',
                  name: 'connectionParams.debug',
                  label: 'Debug',
                  type: 'toggle',
                  helperText: 'If true, enables the debuggin output. Optional, defaults to false.',
                  schema: yup.bool().default(connectionObj.connectionParams?.debug ?? false),
                },
                {
                  component: 'control-select-field',
                  name: 'connectionParams.compress',
                  helperText: 'Algorithm used for compression. Optional, defaults to `none`',
                  label: 'Compression algorithm',
                  options: [...compressValues],
                  schema: yup
                    .string()
                    .oneOf(compressValues)
                    .default(connectionObj.connectionParams?.compress ?? 'none'),
                },
                {
                  component: 'control-number-field',
                  name: 'connectionParams.compress_level',
                  helperText:
                    'Level of compression - with a lower value favoring speed, and a higher value favoring compression. Optional, defaults to 0',
                  label: 'Level of compression',
                  shouldRender: (values: any) =>
                    ['gzip', 'deflate', 'br'].includes(values.connectionParams?.compress),
                  schema: yup
                    .number()
                    .integer()
                    .default(connectionObj.connectionParams?.compress_level ?? 0)
                    .when(['compress'], ([compress], schema) => {
                      if (['gzip', 'deflate'].includes(compress)) {
                        return schema.min(-2).max(9);
                      } else if (compress === 'br') {
                        return schema.min(0).max(11);
                      } else {
                        return schema;
                      }
                    }),
                },
                {
                  component: 'control-number-field',
                  name: 'connectionParams.block_buffer_size',
                  helperText: 'Size of block buffer. Optional, defaults to 2',
                  label: 'Block buffer size',
                  schema: yup
                    .number()
                    .integer()
                    .default(connectionObj.connectionParams?.block_buffer_size ?? 2),
                },
                {
                  component: 'control-number-field',
                  name: 'connectionParams.read_timeout',
                  helperText:
                    'I/O read timeout. Value is in seconds. Optional, defaults to 300 seconds (5 minutes).',
                  label: 'Read Timeout',
                  schema: yup
                    .number()
                    .integer()
                    .default(connectionObj.connectionParams?.read_timeout ?? 300),
                },
                {
                  component: 'control-number-field',
                  name: 'connectionParams.max_compression_buffer',
                  helperText:
                    'Max size of compression buffer during column by column compression. Value is in bytes. Optional, defaults to 1048576 bytes (10MiB).',
                  label: 'Max compression buffer',
                  schema: yup
                    .number()
                    .integer()
                    .default(connectionObj.connectionParams?.max_compression_buffer ?? 1_048_576),
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
                  schema: yup
                    .string()
                    .oneOf(migrateModeValues)
                    .default(initialValues?.migrateMode ?? 'safe')
                    .required(),
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
                  schema: yup
                    .string()
                    .oneOf(writeModeValues)
                    .default(initialValues?.writeMode ?? 'overwrite-delete-stale')
                    .required(),
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
                  schema: yup
                    .number()
                    .integer()
                    .default(initialValues?.spec?.batch_size ?? 10_000)
                    .required(),
                },
                {
                  component: 'control-number-field',
                  name: 'batch_size_bytes',
                  helperText:
                    'Maximum size of items that may be grouped together to be written in a single write. Default is 100,000,000 = 100MB.',
                  label: 'Batch size (bytes)',
                  schema: yup
                    .number()
                    .integer()
                    .default(initialValues?.spec?.batch_size_bytes ?? 100_000_000)
                    .required(),
                },
              ],
            },
          ],
          title: 'Configuration',
        },
      ],
      auth: [AuthType.OTHER],
      guide: {
        title: 'ClickHouse configuration',
        sections: [
          {
            bodies: [
              {
                text: (
                  <>
                    The ClickHouse destination lets you sync data from any CloudQuery source to a
                    ClickHouse&#8209;compatible database.
                  </>
                ),
              },
            ],
          },
          {
            header: 'Step 1: Update ClickHouse Allowlist',
            bodies: [
              {
                text: `To allow CloudQuery network access to your ClickHouse instance, make sure the following
                CloudQuery IPs are in your firewall allowlist:`,
              },
              { code: '35.231.218.115' },
              { code: '35.231.72.234' },
            ],
          },
          {
            shouldRender: (values: any) => values._connectionType === 'fields',
            header: 'Step 2: Gather ClickHouse Cloud connection details',
            bodies: [
              {
                text: (
                  <>
                    1. Click <b>Connect</b> in the ClickHouse sidebar to see the <b>Host</b>, and{' '}
                    <b>Password</b>.
                  </>
                ),
              },
              { image: 'images/connect.png' },
              { image: 'images/cred.png' },
              {
                text: (
                  <>
                    2. Identify the target <b>Database</b> in ClickHouse.
                  </>
                ),
              },
              {
                image: 'images/db.png',
              },
              {
                text: (
                  <>
                    3. Enter the gathered details to the left. Replacing the host port with{' '}
                    <b>9440</b>:
                  </>
                ),
              },
              {
                text: (
                  <>
                    4. Open the <b>Advanced Connection Options</b>, to the left, and toggle{' '}
                    <b>SSL/TLS</b> to be enabled.
                  </>
                ),
              },
            ],
          },
          {
            shouldRender: (values: any) => values._connectionType === 'string',
            header: 'Step 2: Gather ClickHouse Cloud connection details',
            bodies: [
              {
                text: (
                  <>
                    1. Click <b>Connect</b> in the ClickHouse sidebar to see the <b>Host</b>, and{' '}
                    <b>Password</b>.
                  </>
                ),
              },
              { image: 'images/connect.png' },
              { image: 'images/cred.png' },
              {
                text: (
                  <>
                    2. Identify the target <b>Database</b> in ClickHouse.
                  </>
                ),
              },
              {
                image: 'images/db.png',
              },
              {
                text: (
                  <>
                    3. Enter the gathered details interpolated into the <b>Connection String</b>{' '}
                    input to the left. Replacing the host port with <b>9440</b>:
                  </>
                ),
              },
              {
                code: `clickhouse://<USERNAME>:<PASSWORD>@oxhgpaltim.westus3.azure.clickhouse.cloud:9440/<DATABASE>?secure=true`,
              },
            ],
          },
        ],
      },
      errorCodes: {
        INVALID_SPEC: 'Operation failed. The data payload does not meet the required data schema.',
        CONNECTION_FAILED:
          'Failed to establish a connection to ClickHouse. This is rare and might indicate a driver issue.',
        UNAUTHORIZED: 'Operation failed. This might be due to insufficient permissions.',
        UNREACHABLE:
          'The ClickHouse server is unreachable. Check your host, port, and network settings.',
      },
    }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [initialValues],
  );
};
