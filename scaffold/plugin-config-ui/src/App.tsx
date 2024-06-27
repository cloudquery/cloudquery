import { Box, CssBaseline, Stack, ThemeProvider, createTheme } from '@mui/material';
import { Form } from './Form';
import { Fragment, useMemo } from 'react';
import { useFormInit } from './utils/hooks/useFormInit';
import { useWatchPluginUiHeight } from './utils/hooks/useWatchPluginUiHeight';
import { createThemeOptions } from '@cloudquery/cloud-ui';
import { CloudAppMock } from './CloudAppMock';

const DevWrapper = process.env.NODE_ENV === 'development' ? CloudAppMock : Fragment;

function App() {
  const { initialValues, initialized } = useFormInit();
  const containerRef = useWatchPluginUiHeight();

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  return (
    <Box ref={containerRef}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <DevWrapper>
          {initialized && (
            <Stack padding={2}>
              <Form initialValues={initialValues} />
            </Stack>
          )}
        </DevWrapper>
      </ThemeProvider>
    </Box>
  );
}

export default App;
