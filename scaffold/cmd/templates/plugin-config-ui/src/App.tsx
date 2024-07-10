import {
  Box,
  CircularProgress,
  CssBaseline,
  Stack,
  ThemeProvider,
  createTheme,
} from '@mui/material';
import { Form } from './Form';
import { Fragment, useMemo } from 'react';
import { createThemeOptions } from '@cloudquery/cloud-ui';
import { CloudAppMock } from './CloudAppMock';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareInitialValues } from './utils/prepareInitialValues';
import { useFormHeightChange, useFormInit } from '@cloudquery/plugin-config-ui-lib';

const DevWrapper =
  process.env.NODE_ENV === 'production' || window.self !== window.top ? Fragment : CloudAppMock;

function App() {
  const { initialValues, initialized } = useFormInit(pluginUiMessageHandler, false);
  const containerRef = useFormHeightChange(pluginUiMessageHandler);

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  return (
    <Box ref={containerRef}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <DevWrapper>
          {initialized ? (
            <Stack paddingY={2}>
              <Form
                initialValues={initialValues ? prepareInitialValues(initialValues) : undefined}
              />
            </Stack>
          ) : (
            <Stack alignItems="center" paddingY={2}>
              <CircularProgress />
            </Stack>
          )}
        </DevWrapper>
      </ThemeProvider>
    </Box>
  );
}

export default App;
