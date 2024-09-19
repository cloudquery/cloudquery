import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { Link, SourceConfig } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import { pluginUiMessageHandler } from '../utils/messageHandler';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
}

export const useConfig = ({ initialValues }: Props): SourceConfig => {
  return useMemo(
    () => ({
      name: 'hackernews',
      type: 'source',
      label: 'Hacker News',
      docsLink: 'https://hub.cloudquery.io/plugins/source/cloudquery/hackernews/latest/docs',
      iconLink: 'images/logo.webp',
      steps: [
        {
          children: [
            {
              component: 'section',
              title: 'Tables',
              children: [
                {
                  component: 'control-table-selector',
                },
              ],
            },
            {
              component: 'section',
              title: 'Options',
              children: [
                {
                  component: 'control-date-time-field',
                  name: 'start_time',
                  label: 'Start time',
                  clearable: true,
                  helperText:
                    'The earliest news date that the source should fetch. Optional, defaults to no date filtering.',
                  schema: yup
                    .date()
                    .default(
                      initialValues?.spec?.start_time
                        ? new Date(initialValues?.spec?.start_time)
                        : null,
                    )
                    .nullable(),
                },
                {
                  component: 'control-text-field',
                  name: 'item_concurrency',
                  helperText:
                    'Maximum number of news items to fetch concurrently. Recommended value is 100.',
                  label: 'Item concurrency',
                  schema: yup
                    .number()
                    .default(initialValues?.spec?.item_concurrency ?? 100)
                    .required(),
                },
              ],
            },
          ],
          title: 'Configuration',
        },
      ],
      auth: [],
      guide: {
        title: 'Hacker News configuration',
        sections: [
          {
            bodies: [
              {
                text: (
                  <>
                    The Hacker News Source plugin for CloudQuery extracts configuration from the{' '}
                    <Link
                      href="https://github.com/HackerNews/API"
                      pluginUiMessageHandler={pluginUiMessageHandler}
                    >
                      Hacker News API
                    </Link>{' '}
                    and loads it into any supported CloudQuery destination (e.g. PostgreSQL,
                    BigQuery, Snowflake, and{' '}
                    <Link
                      href="https://hub.cloudquery.io/plugins/destination"
                      pluginUiMessageHandler={pluginUiMessageHandler}
                    >
                      more
                    </Link>
                    ).
                  </>
                ),
              },
              {
                text: 'It can be used for real applications, but is mainly intended to serve as an example of CloudQuery Source plugin with an incremental table.',
              },
            ],
          },
        ],
      },
    }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [initialValues],
  );
};
