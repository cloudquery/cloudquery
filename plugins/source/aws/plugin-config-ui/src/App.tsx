import { Box, CssBaseline, Stack, ThemeProvider, createTheme } from '@mui/material';
import { Form } from './Form';
import { Fragment, useMemo } from 'react';
import { createThemeOptions } from '@cloudquery/cloud-ui';
import { CloudAppMock } from './CloudAppMock';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareInitialValues } from './utils/prepareInitialValues';
import { useFormHeightChange, useFormInit } from '@cloudquery/plugin-config-ui-lib';
import { QueryClientProvider } from '@tanstack/react-query';
import { getQueryClient } from './utils/getQueryClient';
import { AuthProvider } from './contexts/authProvider';
const localEnvironment = require('./local-env.json');

const DevWrapper =
  process.env.NODE_ENV === 'production' || window.self !== window.top ? Fragment : CloudAppMock;

const DEV_API_TOKEN =
  process.env.NODE_ENV === 'production' ? undefined : localEnvironment?.DEV_API_TOKEN;

function App() {
  const { initialValues, initialized, apiAuthorizationToken } = useFormInit(
    pluginUiMessageHandler,
    false,
  );

  const containerRef = useFormHeightChange(pluginUiMessageHandler);
  const queryClient = getQueryClient();

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  return (
    <Box ref={containerRef}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <DevWrapper>
          {initialized && (
            <AuthProvider value={DEV_API_TOKEN ?? apiAuthorizationToken}>
              <QueryClientProvider client={queryClient}>
                <Stack paddingY={2}>
                  <Form
                    initialValues={undefined}
                    // initialValues={initialValues ? prepareInitialValues(initialValues) : undefined} // TODO:EDIT
                  />
                </Stack>
              </QueryClientProvider>
            </AuthProvider>
          )}
        </DevWrapper>
      </ThemeProvider>
    </Box>
  );
}

export default App;
