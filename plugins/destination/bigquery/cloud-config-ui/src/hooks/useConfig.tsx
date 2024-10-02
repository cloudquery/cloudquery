import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { AuthType, DestinationConfig } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import { Connect } from '../components/connect';
import { Guide } from '../components/guide';
import { UploadJSON } from '../components/upload';
import { timePartitionOptions } from '../utils/constants';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
}

export const useConfig = ({ initialValues }: Props): DestinationConfig => {
  return useMemo(
    () =>
      ({
        name: 'bigquery',
        type: 'destination',
        label: 'BigQuery',
        docsLink: 'https://hub.cloudquery.io/plugins/destination/cloudquery/bigquery/latest/docs',
        iconLink: 'images/logo.webp',
        stateSchema: {
          _serviceAccount: yup.string().default('Pending launch of the IAM Console...'),
        },
        steps: [
          {
            children: [
              {
                component: 'section',
                title: 'Connect to your database',
                subtitle: 'Set up a connection to your BigQuery instance.',
                children: [
                  {
                    component: 'control-exclusive-toggle',
                    name: '_authType',
                    options: [
                      {
                        label: 'Use CloudQuery Service Account',
                        value: AuthType.OAUTH,
                      },
                      {
                        label: 'Use a service account JSON file',
                        value: AuthType.OTHER,
                      },
                    ],
                    schema: yup
                      .mixed()
                      .oneOf(Object.values(AuthType))
                      .default(
                        initialValues?.spec?.service_account_key_json
                          ? AuthType.OTHER
                          : AuthType.OAUTH,
                      ),
                  },
                  {
                    component: 'sub-section',
                    shouldRender: (values: any) => values._authType === AuthType.OAUTH,
                    children: [Connect],
                  },
                  {
                    component: 'sub-section',
                    shouldRender: (values: any) => values._authType === AuthType.OTHER,
                    children: [
                      {
                        name: 'service_account_key_json',
                        component: UploadJSON,
                        schema: yup
                          .string()
                          .default(initialValues?.spec?.service_account_key_json ?? '')
                          .when('_authType', {
                            is: (_authType: AuthType) => _authType === AuthType.OTHER,
                            // eslint-disable-next-line unicorn/no-thenable
                            then: (schema: any) => schema.required(),
                          }),
                      },
                    ],
                  },
                ],
              },
              {
                component: 'section',
                title: 'Connection Options',
                children: [
                  {
                    component: 'control-secret-field',
                    name: 'project_id',
                    helperText:
                      'The id of the project where the destination BigQuery database resides.',
                    label: 'Google Cloud Project ID',
                    schema: yup
                      .string()
                      .default(initialValues?.spec?.project_id ?? '')
                      .required(),
                  },
                  {
                    component: 'control-secret-field',
                    name: 'dataset_id',
                    helperText:
                      'The name of the BigQuery dataset within the project, e.g. my_dataset. This dataset needs to be created before running a sync or migration.',
                    label: 'Google Cloud BigQuery DataSet ID',
                    schema: yup
                      .string()
                      .default(initialValues?.spec?.dataset_id ?? '')
                      .required(),
                  },
                ],
              },
              {
                component: 'collapsible-section',
                title: 'Advanced Connection Options',
                defaultExpanded: false,
                children: [
                  {
                    component: 'control-text-field',
                    name: 'dataset_location',
                    helperText:
                      'The data location of the BigQuery dataset. If set, will be used as the default location for job operations.',
                    label: 'Google Cloud BigQuery DataSet Location',
                    schema: yup.string().default(initialValues?.spec?.dataset_location ?? ''),
                  },
                  {
                    component: 'control-select-field',
                    name: 'time_partitioning',
                    options: [...timePartitionOptions],
                    helperText: (
                      <>
                        The time partitioning to use when creating tables. The partition time column
                        used will always be <b>_cq_sync_time</b> so that all rows for a sync run
                        will be partitioned on the hour/day the sync started.
                      </>
                    ),
                    label: 'Time Partitioning',
                    schema: yup
                      .string()
                      .oneOf(timePartitionOptions)
                      .default(initialValues?.spec?.time_partitioning ?? 'none'),
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
        guide: Guide,
      }) as DestinationConfig,
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [initialValues],
  );
};
