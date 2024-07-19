import Stack from '@mui/material/Stack';
import CssBaseline from '@mui/material/CssBaseline';
import ThemeProvider from '@mui/material/styles/ThemeProvider';
import createTheme from '@mui/material/styles/createTheme';
import { Form } from './form';
import { Fragment, useMemo } from 'react';
import { createThemeOptions } from '@cloudquery/cloud-ui';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareInitialValues } from './utils/prepareInitialValues';
import { CloudAppMock, useFormHeightChange, useFormInit } from '@cloudquery/plugin-config-ui-lib';

const useCloudAppMock = process.env.NODE_ENV !== 'production' && window.self === window.top;
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
          <Stack>
            <Form initialValues={initialValues ? prepareInitialValues(initialValues) : undefined} />
          </Stack>
        )}
      </DevWrapper>
    </ThemeProvider>
  );
}

export default App;
