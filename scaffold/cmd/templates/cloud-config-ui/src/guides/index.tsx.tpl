import { RenderGuide, SetupGuide } from '@cloudquery/plugin-config-ui-lib';
import Stack from '@mui/material/Stack';

import { pluginUiMessageHandler } from '../utils/messageHandler';

export function Guides() {
  return (
    <SetupGuide
      title="{{.Name}} config"
      pluginUiMessageHandler={pluginUiMessageHandler}
      docsLink="https://hub.cloudquery.io/plugins/source/cloudquery/{{.Name}}/latest/docs"
    >
      <Stack spacing={3}>
        <RenderGuide
          pluginUiMessageHandler={pluginUiMessageHandler}
          sections={[
            {
              bodies: [
                {
                  text: (
                    <>
                      {`The {{.Name}} source lets you sync data to any
                      CloudQuery destination.`}
                    </>
                  ),
                },
              ],
            },
          ]}
        />
        <RenderGuide
          pluginUiMessageHandler={pluginUiMessageHandler}
          sections={[
            {
              header: `Setup guide`,
              bodies: [
                {
                  text: `1. Enter a Name for your plugin.`,
                },
                {
                  text: `2. Enter a Token for your plugin.`,
                },
                {
                  text: `3. Select tables to sync data from.`,
                },
              ],
            },
          ]}
        />
      </Stack>
    </SetupGuide>
  );
}
