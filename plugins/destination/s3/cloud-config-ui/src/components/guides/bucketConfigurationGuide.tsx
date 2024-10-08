import { Link, RenderGuide } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

const SECTIONS = [
  {
    header: 'Bucket configuration',
    bodies: [
      {
        text: (
          <>
            Specify the region, path, and a format to sync in. Path can be configured using macros.
            For the list of macros, see the detailed plugin{' '}
            <Link
              href="https://hub.cloudquery.io/plugins/destination/cloudquery/s3/latest/docs"
              pluginUiMessageHandler={pluginUiMessageHandler}
            >
              documentation
            </Link>
            .
          </>
        ),
      },
    ],
  },
  {
    header: 'Format spec',
    bodies: [
      {
        text: 'Depending on the format you choose, you can specify additional options for how the data will be written to S3.',
      },
    ],
  },
  {
    header: 'Server side encryption',
    bodies: [
      {
        text: 'You can enable server side encryption for the data written to S3 along with its options.',
      },
    ],
  },
  {
    header: 'Advanced options',
    bodies: [
      {
        text: 'You can also specify additional options for how the data will be written to S3. They depend on the format you choose.',
      },
    ],
  },
];

export function BucketConfigurationGuide() {
  return <RenderGuide pluginUiMessageHandler={pluginUiMessageHandler} sections={SECTIONS} />;
}
