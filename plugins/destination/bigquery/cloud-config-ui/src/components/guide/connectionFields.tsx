import { GuideSectionBody, Link } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

export const connectionFields: GuideSectionBody[] = [
  {
    text: (
      <>
        1. Open the{' '}
        <Link
          pluginUiMessageHandler={pluginUiMessageHandler}
          href="https://console.cloud.google.com/bigquery"
        >
          BigQuery console
        </Link>
      </>
    ),
  },
  {
    text: (
      <>
        2. Locate the <b>Google Cloud Project ID</b> for your desired dataset, and enter it on the
        left.
      </>
    ),
  },
  { image: 'images/conn-1.webp' },
  {
    text: (
      <>
        3. Locate the <b>Google Cloud BigQuery DataSet ID</b> for your desired dataset, and enter it
        on the left.
      </>
    ),
  },
  { image: 'images/conn-2.webp' },
];
