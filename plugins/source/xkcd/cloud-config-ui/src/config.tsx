import { Link, SourceConfig } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from './utils/messageHandler';

const config: SourceConfig = {
  name: 'xkcd',
  type: 'source',
  label: 'XKCD',
  docsLink: 'https://hub.cloudquery.io/plugins/source/cloudquery/xkcd/latest/docs',
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
      ],
      title: 'Configuration',
    },
  ],
  auth: [],
  guide: {
    title: 'XKCD configuration',
    sections: [
      {
        bodies: [
          {
            text: (
              <>
                This CloudQuery source plugin fetches data from the{' '}
                <Link
                  href="https://xkcd.com/json.html"
                  pluginUiMessageHandler={pluginUiMessageHandler}
                >
                  XKCD API
                </Link>
                , allowing you to load the XKCD comic data into any CloudQuery-supported destination
                (e.g. PostgreSQL, Elasticsearch, Snowflake, etc.). See{' '}
                <Link
                  href="https://www.cloudquery.io/docs/plugins/destinations/overview"
                  pluginUiMessageHandler={pluginUiMessageHandler}
                >
                  CloudQuery destinations
                </Link>{' '}
                for a complete list of supported destinations.
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
