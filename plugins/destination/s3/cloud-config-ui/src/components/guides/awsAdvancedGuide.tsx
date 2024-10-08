import { RenderGuide } from '@cloudquery/plugin-config-ui-lib';

import { pluginUiMessageHandler } from '../../utils/messageHandler';

export function AWSAdvancedGuide() {
  return (
    <RenderGuide
      pluginUiMessageHandler={pluginUiMessageHandler}
      sections={[
        {
          bodies: [
            {
              text: 'This set of advanced options helps you improving sync performance and dealing with service rate limits. If you have not run a single sync, we suggest you keep the recommended defaults.',
            },
          ],
        },
      ]}
    />
  );
}
