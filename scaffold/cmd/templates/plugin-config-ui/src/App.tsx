import { Box, CssBaseline, Stack, ThemeProvider, createTheme } from '@mui/material';
import { Form } from './Form';
import { Fragment, useMemo } from 'react';
import {
  createThemeOptions,
  usePluginUiFormHeightChange,
  usePluginUiFormInit,
} from '@cloudquery/cloud-ui';
import { CloudAppMock } from './CloudAppMock';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareInitialValues } from './utils/prepareInitialValues';

const DevWrapper =
  process.env.NODE_ENV === 'production' || window.self !== window.top ? Fragment : CloudAppMock;

function App() {
  const { initialValues, initialized } = usePluginUiFormInit(pluginUiMessageHandler, false);
  const containerRef = usePluginUiFormHeightChange(pluginUiMessageHandler);

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  return (
    <Box ref={containerRef}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <DevWrapper>
          {initialized && (
            <Stack paddingY={2}>
              <Form
                initialValues={initialValues ? prepareInitialValues(initialValues) : undefined}
              />
            </Stack>
          )}
        </DevWrapper>
      </ThemeProvider>
    </Box>
  );
}

export default App;
