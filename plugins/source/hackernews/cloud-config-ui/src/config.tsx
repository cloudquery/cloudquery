import { Link, SourceConfig } from '@cloudquery/plugin-config-ui-lib';

import { StartTime } from './form/components/startTime';
import { pluginUiMessageHandler } from './utils/messageHandler';

const config: SourceConfig = {
  name: 'hackernews',
  type: 'source',
  label: 'Hacker News',
  docsLink: 'https://hub.cloudquery.io/plugins/source/cloudquery/hackernews/latest/docs',
  iconLink: 'images/logo.webp',
  steps: [
    {
      sections: [
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
            StartTime,
            {
              component: 'control-text-field',
              name: 'item_concurrency',
              helperText:
                'Maximum number of news items to fetch concurrently. Recommended value is 100.',
              label: 'Item concurrency',
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
                and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery,
                Snowflake, and{' '}
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
};

export default config;
