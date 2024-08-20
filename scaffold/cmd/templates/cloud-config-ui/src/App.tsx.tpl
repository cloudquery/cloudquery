import { Fragment, useMemo } from 'react';

import { createThemeOptions } from '@cloudquery/cloud-ui';
import { CloudAppMock, useFormHeightChange, useFormInit } from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Stack from '@mui/material/Stack';
import createTheme from '@mui/material/styles/createTheme';
import ThemeProvider from '@mui/material/styles/ThemeProvider';

import { Form } from './form';

import { Guides } from './guides';
import { pluginUiMessageHandler } from './utils/messageHandler';

const useCloudAppMock =
  (process.env.REACT_APP_USE_CLOUD_APP_MOCK === 'true' || process.env.NODE_ENV !== 'production') &&
  window.self === window.top;
const DevWrapper = useCloudAppMock ? CloudAppMock : Fragment;
// eslint-disable-next-line unicorn/prefer-module
const devWrapperProps: any = useCloudAppMock ? require('./.env.json') : undefined;

function App() {
  const { initialValues, initialized } = useFormInit(pluginUiMessageHandler, false);
  useFormHeightChange(pluginUiMessageHandler);

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <DevWrapper {...devWrapperProps}>
        {initialized && (
          <Stack direction="row" gap={3} flexWrap="wrap">
            <Box flex="1 1 0" minWidth={480}>
              <Form initialValues={initialValues} />
            </Box>
            <Box width={360} minWidth={360}>
              <Guides />
            </Box>
          </Stack>
        )}
      </DevWrapper>
    </ThemeProvider>
  );
}

export default App;
