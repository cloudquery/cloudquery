import { RenderGuide, SetupGuide } from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Stack from '@mui/material/Stack';
import useTheme from '@mui/material/styles/useTheme';

import { pluginUiMessageHandler } from '../utils/messageHandler';

export function Guides() {
  const { palette } = useTheme();

  return (
    <SetupGuide
      title="MySQL config"
      pluginUiMessageHandler={pluginUiMessageHandler}
      docsLink="https://hub.cloudquery.io/plugins/destination/cloudquery/mysql/latest/docs"
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
                      The MySQL destination lets you sync data from any CloudQuery source to a
                      MySQL&#8209;compatible database.
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
              header: 'Setup guide',
              bodies: [
                {
                  text: (
                    <>
                      {/* TODO : These IPs relevant for mysql plugin? */}
                      <Box>
                        To allow CloudQuery network access to your MySQL instance, make sure the
                        following CloudQuery IPs are in your firewall allowlist:
                      </Box>
                      <List
                        sx={{
                          listStyleType: 'disc',
                          listStylePosition: 'inside',
                        }}
                      >
                        <ListItem
                          sx={{
                            paddingLeft: 0.5,
                            paddingY: 0.5,
                            display: 'list-item',
                            '&::marker': { color: palette.text.secondary },
                          }}
                        >
                          <Box component="span" color={palette.text.secondary}>
                            35.231.218.115
                          </Box>
                        </ListItem>
                        <ListItem
                          sx={{
                            paddingLeft: 0.5,
                            paddingY: 0.5,
                            display: 'list-item',
                            '&::marker': { color: palette.text.secondary },
                          }}
                        >
                          <Box component="span" color={palette.text.secondary}>
                            35.231.72.234
                          </Box>
                        </ListItem>
                      </List>
                    </>
                  ),
                },
              ],
            },
          ]}
        />
      </Stack>
    </SetupGuide>
  );
}
