import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { AuthType, DestinationConfig, Link } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import {
  connectionTypeValues,
  migrateModeValues,
  pgxLogLevelValues,
  sslModeValues,
  writeModeValues,
} from '../utils/constants';
import { convertConnectionStringToFields } from '../utils/convertConnectionStringToFields';
import { pluginUiMessageHandler } from '../utils/messageHandler';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
  isManagedDestination: boolean;
}

export const useConfig = ({ initialValues, isManagedDestination }: Props): DestinationConfig => {
  const url = initialValues?.spec?.connection_string || '';
  const connectionObj: Record<string, any> = useMemo(
    () => convertConnectionStringToFields(url),
    [url],
  );

  return useMemo(
    () => ({
      name: 'postgresql',
      type: 'destination',
      label: 'PostgreSQL',
      docsLink: 'https://hub.cloudquery.io/plugins/source/cloudquery/postgresql/latest/docs',
      iconLink: 'images/postgresql.png',
      steps: [
        {
          children: [
            {
              component: 'section',
              title: 'Connect to your database',
              subtitle: 'Set up a connection to your PostgreSQL instance.',
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
                    .default(
                      isManagedDestination || url.startsWith('postgresql://') ? 'string' : 'fields',
                    )
                    .required(),
                },
                {
                  component: 'control-secret-field',
                  name: 'connection_string',
                  helperText:
                    'Connection string to connect to the database. E.g. postgres://user:pass@localhost:5432/mydb?sslmode=prefer',
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
                  component: 'control-text-field',
                  name: 'host',
                  helperText: 'Host to connect to. E.g. 1.2.3.4 or mydb.host.com.',
                  label: 'Host',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(253)
                    .default(connectionObj.host ?? '')
                    .when('_connectionType', {
                      is: 'fields',
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-text-field',
                  name: 'port',
                  helperText: 'Port to connect to. Optional, defaults to 5432.',
                  label: 'Port',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(5)
                    .matches(/^($)|(\d+)$/, 'Port must be a number')
                    .default(connectionObj.port ?? ''),
                },
                {
                  component: 'control-text-field',
                  name: 'database',
                  helperText: 'Name of the PostgreSQL database you want to connect to.',
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
                  component: 'control-text-field',
                  name: 'user',
                  helperText: 'Username to use when authenticating. Optional, defaults to empty.',
                  label: 'Username',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(63)
                    .default(connectionObj.user ?? ''),
                },
                {
                  component: 'control-secret-field',
                  name: 'password',
                  label: 'Password',
                  helperText: 'Password to use when authenticating. Optional, defaults to empty.',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup
                    .string()
                    .max(63)
                    .default(connectionObj.password ?? ''),
                },
                {
                  component: 'control-text-field',
                  name: 'connectionParams.search_path',
                  helperText:
                    'Name of the PostgreSQL schema you want to connect to. Optional, defaults to public.',
                  label: 'Schema',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup.string().default(connectionObj.connectionParams?.search_path ?? ''),
                },
                {
                  component: 'control-boolean-field',
                  name: 'connectionParams.ssl',
                  label: 'SSL',
                  type: 'toggle',
                  shouldRender: (values: any) => values._connectionType === 'fields',
                  schema: yup.bool().default(connectionObj.connectionParams?.ssl ?? false),
                },
                {
                  component: 'control-select-field',
                  name: 'connectionParams.sslmode',
                  helperText:
                    'SSL connections to encrypt client/server communications using TLS protocols for increased security.',
                  label: 'SSL Mode',
                  shouldRender: (values: any) =>
                    values._connectionType === 'fields' && !!values.connectionParams?.ssl,
                  options: [...sslModeValues],
                  schema: yup
                    .string()
                    .oneOf(sslModeValues)
                    .default(connectionObj.connectionParams?.sslmode ?? 'require'),
                },
              ],
            },
            {
              component: 'section',
              title: 'Sync Options',
              subtitle: 'Configure how CloudQuery should write to your destination.',
              children: [
                {
                  component: 'control-select-field',
                  name: 'migrateMode',
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
                  label: 'Migrate Mode',
                  options: [...migrateModeValues],
                  schema: yup
                    .string()
                    .oneOf(migrateModeValues)
                    .default(initialValues?.migrateMode ?? 'safe')
                    .required(),
                },
                {
                  component: 'control-select-field',
                  name: 'writeMode',
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
                  label: 'Write Mode',
                  options: [...writeModeValues],
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
              title: 'Advanced Options',
              defaultExpanded: false,
              children: [
                {
                  component: 'control-select-field',
                  name: 'pgx_log_level',
                  helperText: 'Configure the level of detail of the log from this destination.',
                  label: 'Log level',
                  options: [...pgxLogLevelValues],
                  schema: yup
                    .string()
                    .oneOf(pgxLogLevelValues)
                    .default(initialValues?.spec?.pgx_log_level ?? 'error')
                    .required(),
                },
                {
                  component: 'control-text-field',
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
                  component: 'control-text-field',
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
                {
                  component: 'control-text-field',
                  name: 'batch_timeout',
                  helperText: 'Maximum interval between batch writes. Defaults to 60s.',
                  label: 'Batch timeout',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.batch_timeout ?? '60s')
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
        title: 'PostgreSQL configuration',
        sections: [
          {
            bodies: [
              {
                text: (
                  <>
                    The PostgreSQL destination lets you sync data from any CloudQuery source to a
                    PostgreSQL‑compatible database.
                  </>
                ),
              },
            ],
          },
          {
            header: 'Setup guide',
            bodies: [
              {
                text: `To allow CloudQuery network access to your PostgreSQL instance, make sure the following
                CloudQuery IPs are in your firewall allowlist:`,
              },
              { code: '35.231.218.115' },
              { code: '35.231.72.234' },
            ],
          },
        ],
      },
      errorCodes: {
        INVALID_CONFIG:
          'The connection string (DSN) is invalid or in an incorrect format. Please check and correct your connection details.',
        UNKNOWN_DATABASE:
          "The specified database does not exist. Please check your database name and ensure it's created on the server.",
        DNS_FAILED:
          "Failed to resolve the PostgreSQL host. Please ensure you've specified a valid host and that it's reachable.",
        AUTH_FAILED:
          'Failed to authenticate with the PostgreSQL server. Please check your username and password.',
        CONN_FAILED:
          'Failed to establish a connection to the PostgreSQL database. This is rare and might indicate a driver issue.',
        UNKNOWN_SCHEMA:
          "The specified schema does not exist. Please check your schema name and ensure it's created on the server.",
      },
    }),
    [connectionObj, initialValues, isManagedDestination, url],
  );
};
