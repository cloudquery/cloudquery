import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { AuthType, DestinationConfig, Link } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import { Guides } from '../components/guides';
import { OAuthConnect } from '../components/oauthConnect';
import { authSubmitGuard } from '../utils/authSubmitGuard';
import { pluginUiMessageHandler } from '../utils/messageHandler';

export enum SetupType {
  Manual = 'manual',
  Console = 'console',
}
export const setupTypes = Object.values(SetupType);

export const useConfig = ({
  initialValues,
}: {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
}): DestinationConfig => {
  return useMemo(
    () => ({
      name: 's3',
      type: 'destination',
      label: 'S3',
      docsLink: 'https://hub.cloudquery.io/plugins/destination/cloudquery/s3/latest/docs',
      iconLink: 'images/logo.webp',
      steps: [
        {
          title: 'Connect to your S3 bucket',
          children: [
            {
              component: 'section',
              title: 'Authentication',
              subtitle: `Set up a connection to your S3 bucket`,
              children: [
                {
                  title: 'Authentication',
                  component: 'control-exclusive-toggle',
                  name: '_authType',
                  options: [
                    {
                      label: 'AWS Console',
                      value: AuthType.OAUTH,
                    },
                    {
                      label: 'Manual setup',
                      value: AuthType.OTHER,
                    },
                  ],
                  schema: yup.mixed().oneOf(Object.values(AuthType)).default(AuthType.OAUTH),
                },
                {
                  component: 'sub-section',
                  shouldRender: (values: any) => values._authType === AuthType.OAUTH,
                  children: [OAuthConnect.bind({}, { pluginUiMessageHandler })],
                },
                {
                  component: 'sub-section',
                  shouldRender: (values: any) => values._authType === AuthType.OTHER,
                  children: [
                    {
                      component: 'control-secret-field',
                      name: 'arn',
                      helperText:
                        'Amazon Resource Name uniquely identifies AWS resources. It will be provided when you finish creating the new role',
                      label: 'ARN',
                      schema: yup
                        .string()
                        .default(initialValues?.spec?.arn ?? '')
                        .trim()
                        .matches(
                          new RegExp(
                            /^arn:(aws|aws-us-gov|aws-cn):iam:(\w+(?:-\w+)+)?:\d{12}:role\/[\w.-]+$/,
                          ),
                          {
                            message:
                              'ARN must be a valid AWS ARN format, example: arn:aws:iam::123456789012:role/CloudQueryIntegrationRoleForAWSSource',
                            name: 'arn',
                          },
                        )
                        .required('ARN cannot be empty'),
                    },
                    {
                      component: 'control-secret-field',
                      name: 'externalId',
                      label: 'External ID',
                      helperText:
                        'External ID must match the connected AWS Role Trust Policy. See documentation on the right for more information.',
                      schema: yup
                        .string()
                        .default(initialValues?.spec?.externalId ?? '')
                        .trim()
                        .when('_authType', {
                          is: AuthType.OTHER,
                          // eslint-disable-next-line unicorn/no-thenable
                          then: (schema) => schema.required('External ID is required'),
                        }),
                    },
                    {
                      component: 'control-text-field',
                      name: 'bucket',
                      helperText: 'Name of the S3 bucket.',
                      label: 'Bucket *',
                      schema: yup
                        .string()
                        .default(initialValues?.spec?.bucket ?? '')
                        .trim()
                        .required(),
                    },
                  ],
                },
              ],
            },
          ],
        },
        {
          submitGuard: authSubmitGuard,
          children: [
            {
              component: 'section',
              title: 'Bucket Configuration',
              subtitle: 'Configure the bucket options.',
              children: [
                {
                  component: 'control-text-field',
                  name: 'region',
                  helperText: 'Region where the bucket is located. E.g. us-east-1.',
                  label: 'Region *',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.region ?? '')
                    .trim()
                    .when('_step', {
                      is: 1,
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-text-field',
                  name: 'path',
                  helperText: 'Path to where the files will be uploaded in the bucket.',
                  label: 'Path *',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.path ?? '')
                    .trim()
                    .when('_step', {
                      is: 1,
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) =>
                        schema
                          .required()
                          .test(
                            'clean-path',
                            'value must not contain ./ or //',
                            (value) => !/(\.\/|\/\/)/.test(value ?? ''),
                          )
                          .when('no_rotate', {
                            is: true,
                            // eslint-disable-next-line unicorn/no-thenable
                            then: (schema) =>
                              schema.test(
                                'no-uuid-in-path',
                                'the {{UUID}} placeholder must not be present in the path when no_rotate is enabled',
                                (value) => !/{{UUID}}/.test(value ?? ''),
                              ),
                            otherwise: (schema) =>
                              schema.test(
                                'uuid-in-path',
                                'the {{UUID}} placeholder must be present in the path',
                                (value) => /{{UUID}}/.test(value ?? ''),
                              ),
                          }),
                    }),
                },
                {
                  component: 'control-select-field',
                  name: 'format',
                  helperText:
                    'Format of the output file. Supported values are csv, json, and parquet.',
                  label: 'Format *',
                  options: ['csv', 'json', 'parquet'],
                  schema: yup
                    .string()
                    .oneOf(['csv', 'json', 'parquet'])
                    .default(initialValues?.spec?.format ?? 'csv')
                    .when('_step', {
                      is: 1,
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
              ],
            },
            {
              component: 'collapsible-section',
              defaultExpanded: false,
              title: 'Format Spec',
              shouldRender: (values) => ['csv', 'parquet'].includes(values.format),
              children: [
                {
                  component: 'control-text-field',
                  name: 'format_spec_csv_delimiter',
                  shouldRender: (values) => values.format === 'csv',
                  helperText: 'Delimiter to use in the CSV file.',
                  label: 'Delimiter',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.format_spec?.delimiter ?? ',')
                    .trim(),
                },
                {
                  component: 'control-boolean-field',
                  name: 'format_spec_csv_skip_header',
                  shouldRender: (values) => values.format === 'csv',
                  helperText:
                    'If set to true, the CSV file will not contain a header row as the first row.',
                  label: 'Skip Header',
                  type: 'toggle',
                  schema: yup
                    .boolean()
                    .default(
                      typeof initialValues?.spec?.format_spec?.skip_header === 'boolean'
                        ? initialValues?.spec?.format_spec?.skip_header
                        : false,
                    ),
                },
                {
                  component: 'control-select-field',
                  options: ['v1.0', 'v2.4', 'v2.6', 'v2Latest'],
                  name: 'format_spec_parquet_version',
                  shouldRender: (values) => values.format === 'parquet',
                  helperText: `Supported values are v1.0, v2.4, v2.6 and v2Latest.\nv2Latest is an alias for the latest version available in the Parquet library which is currently v2.6.nUseful when the reader consuming the Parquet files does not support the latest version.`,
                  label: 'Parquet format version to use',
                  schema: yup
                    .string()
                    .oneOf(['v1.0', 'v2.4', 'v2.6', 'v2Latest'])
                    .default(initialValues?.spec?.format_spec?.version ?? 'v2Latest'),
                },
                {
                  component: 'control-text-field',
                  name: 'format_spec_parquet_root_repetition',
                  shouldRender: (values) => values.format === 'parquet',
                  helperText: (
                    <>
                      Supported values are undefined, required, optional and repeated.{' '}
                      <Link
                        pluginUiMessageHandler={pluginUiMessageHandler}
                        href="https://github.com/apache/arrow/issues/20243"
                      >
                        Learn more
                      </Link>
                    </>
                  ),
                  label: 'Root node repetition type',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.format_spec?.root_repetition ?? 'repeated')
                    .trim(),
                },
              ],
            },
            {
              component: 'collapsible-section',
              title: 'Server-side encryption',
              defaultExpanded: false,
              children: [
                {
                  component: 'control-boolean-field',
                  name: 'server_side_encryption_configuration_active',
                  label: 'Enable server-side encryption',
                  type: 'toggle',
                  schema: yup
                    .boolean()
                    .default(!!initialValues?.spec?.server_side_encryption_configuration),
                },
                {
                  component: 'control-text-field',
                  name: 'server_side_encryption_configuration_sse_kms_key_id',
                  shouldRender: (values) =>
                    values.server_side_encryption_configuration_active === true,
                  helperText:
                    'KMS Key ID appended to S3 API calls header. Used in conjunction with server_side_encryption.',
                  label: 'KMS Key ID',
                  schema: yup
                    .string()
                    .default(
                      initialValues?.spec?.server_side_encryption_configuration?.sse_kms_key_id ??
                        '',
                    )
                    .trim()
                    .when(['server_side_encryption_configuration_active', '_step'], {
                      is: (active: boolean, step: number) => active === true && step === 1,
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
                {
                  component: 'control-text-field',
                  name: 'server_side_encryption_configuration_server_side_encryption',
                  shouldRender: (values) =>
                    values.server_side_encryption_configuration_active === true,
                  helperText:
                    'The server-side encryption algorithm used when storing the object in S3. Supported values are AES256, aws:kms and aws:kms:dsse.',
                  label: 'Server-side encryption algorithm',
                  schema: yup
                    .string()
                    .default(
                      initialValues?.spec?.server_side_encryption_configuration
                        ?.server_side_encryption ?? '',
                    )
                    .trim()
                    .when(['server_side_encryption_configuration_active', '_step'], {
                      is: (active: boolean, step: number) => active === true && step === 1,
                      // eslint-disable-next-line unicorn/no-thenable
                      then: (schema) => schema.required(),
                    }),
                },
              ],
            },
            {
              component: 'collapsible-section',
              title: 'Advanced options',
              defaultExpanded: false,
              children: [
                {
                  shouldRender: (values) => values.format !== 'parquet',
                  component: 'control-select-field',
                  name: 'compression',
                  helperText: 'Compression algorithm to use.',
                  label: 'Compression',
                  options: [
                    {
                      label: '(empty)',
                      value: '',
                    },
                    'gzip',
                  ],
                  schema: yup
                    .string()
                    .oneOf(['', 'gzip'])
                    .default(initialValues?.spec?.compression ?? ''),
                },
                {
                  component: 'control-boolean-field',
                  name: 'no_rotate',
                  label: 'No Rotate',
                  type: 'toggle',
                  helperText:
                    'If set to true, the plugin will write to one file per table. Otherwise, for every batch a new file will be created with a different .<UUID> suffix.',
                  schema: yup.bool().default(!!initialValues?.spec?.no_rotate || false),
                },
                {
                  component: 'control-boolean-field',
                  name: 'athena',
                  label: 'Athena',
                  type: 'toggle',
                  helperText:
                    'When athena is set to true, the S3 plugin will sanitize keys in JSON columns to be compatible with the Hive Metastore / Athena.\nThis allows tables to be created with a Glue Crawler and then queried via Athena, without changes to the table schema.',
                  schema: yup.bool().default(!!initialValues?.spec?.athena || false),
                },
                {
                  shouldRender: (values) => values.format === 'parquet',
                  component: 'control-boolean-field',
                  name: 'write_empty_objects_for_empty_tables',
                  label: 'Write Empty Objects for Empty Tables',
                  type: 'toggle',
                  helperText:
                    "By default only tables with resources are persisted to objects during the sync. If you'd like to persist empty objects for empty tables enable this option. Useful when using CloudQuery Compliance policies to ensure all tables have their schema populated by a query engine like Athena.",
                  schema: yup
                    .bool()
                    .default(!!initialValues?.spec?.write_empty_objects_for_empty_tables || false),
                },
                {
                  component: 'control-boolean-field',
                  name: 'test_write',
                  label: 'Test Write',
                  type: 'toggle',
                  helperText:
                    'Ensure write access to the given bucket and path by writing a test object on each sync.\nIf you are sure that the bucket and path are writable, you can set this to false to skip the test.',
                  schema: yup.bool().default(!!initialValues?.spec?.test_write || true),
                },
                {
                  component: 'control-text-field',
                  name: 'endpoint',
                  helperText:
                    'Endpoint to use for S3 API calls. This is useful for S3-compatible storage services such as MinIO.\nNote: if you want to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY, use_path_style should be enabled, too.',
                  label: 'Endpoint',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.endpoint ?? '')
                    .trim(),
                },
                {
                  component: 'control-select-field',
                  options: [
                    {
                      label: '(empty)',
                      value: '',
                    },
                    'private',
                    'public-read',
                    'public-read-write',
                    'authenticated-read',
                    'aws-exec-read',
                    'bucket-owner-read',
                    'bucket-owner-full-control',
                  ],
                  name: 'acl',
                  helperText:
                    'Canned ACL to apply to the object. Supported values are private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, bucket-owner-full-control.',
                  label: 'Canned ACL',
                  schema: yup
                    .string()
                    .oneOf([
                      '',
                      'private',
                      'public-read',
                      'public-read-write',
                      'authenticated-read',
                      'aws-exec-read',
                      'bucket-owner-read',
                      'bucket-owner-full-control',
                    ])
                    .default(initialValues?.spec?.acl ?? ''),
                },
                {
                  component: 'control-boolean-field',
                  name: 'endpoint_skip_tls_verify',
                  label: 'Endpoint Skip TLS Verify',
                  type: 'toggle',
                  helperText:
                    'Disable TLS verification for requests to your S3 endpoint.\nThis option is intended to be used when using a custom endpoint using the endpoint option.',
                  schema: yup
                    .bool()
                    .default(!!initialValues?.spec?.endpoint_skip_tls_verify || false),
                },
                {
                  component: 'control-boolean-field',
                  name: 'use_path_style',
                  label: 'Use Path Style',
                  type: 'toggle',
                  helperText:
                    'Allows to use path-style addressing in the endpoint option, i.e., https://s3.amazonaws.com/BUCKET/KEY.\nBy default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).',
                  schema: yup.bool().default(!!initialValues?.spec?.use_path_style || false),
                },
                {
                  shouldRender: (values) => values.no_rotate === false,
                  component: 'control-number-field',
                  name: 'batch_size',
                  helperText: 'Number of records to write before starting a new object.',
                  label: 'Batch Size',
                  schema: yup
                    .number()
                    .integer()
                    .positive()
                    .default(initialValues?.spec?.batch_size ?? 10_000),
                },
                {
                  shouldRender: (values) => values.no_rotate === false,
                  component: 'control-number-field',
                  name: 'batch_size_bytes',
                  helperText: 'Number of bytes to write before starting a new object.',
                  label: 'Batch Size (Bytes)',
                  schema: yup
                    .number()
                    .integer()
                    .positive()
                    .default(initialValues?.spec?.batch_size_bytes ?? 52_428_800),
                },
                {
                  shouldRender: (values) => values.no_rotate === false,
                  component: 'control-text-field',
                  name: 'batch_timeout',
                  helperText: 'Maximum interval between batch writes.',
                  label: 'Batch Timeout',
                  schema: yup
                    .string()
                    .default(initialValues?.spec?.batch_timeout ?? '30s')
                    .trim(),
                },
              ],
            },
          ],
          title: 'Configuration',
        },
      ],
      auth: [AuthType.OTHER],
      guide: Guides,
      stateSchema: {
        // those fields are used to store connector authentication state in order to prevent
        // reauthentication when it's not necessary.
        _finishedConnectorId: yup.string().default(initialValues?.connectorId ?? ''),
        _finishedExternalId: yup.string().default(initialValues?.spec?.externalId ?? ''),
        _finishedArn: yup.string().default(initialValues?.spec?.arn ?? ''),
      },
    }),
    [initialValues],
  );
};
