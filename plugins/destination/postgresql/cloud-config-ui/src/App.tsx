import CssBaseline from '@mui/material/CssBaseline';
import ThemeProvider from '@mui/material/styles/ThemeProvider';
import createTheme from '@mui/material/styles/createTheme';
import { Fragment, useMemo } from 'react';
import { createThemeOptions } from '@cloudquery/cloud-ui';
import { pluginUiMessageHandler } from './utils/messageHandler';
import {
  CloudAppMock,
  ConfigUIForm,
  // useFormHeightChange,
  useFormInit,
  PluginContextProvider,
} from '@cloudquery/plugin-config-ui-lib';
import { useConfig } from './hooks/useConfig';
import Box from '@mui/material/Box';
import { prepareSubmitValues } from './utils/prepareSubmitValues';

const useCloudAppMock =
  (process.env.REACT_APP_USE_CLOUD_APP_MOCK === 'true' || process.env.NODE_ENV !== 'production') &&
  window.self === window.top;
const DevWrapper = useCloudAppMock ? CloudAppMock : Fragment;
// eslint-disable-next-line unicorn/prefer-module
const { plugin, ...devWrapperProps }: any = useCloudAppMock ? require('./.env.json') : {};

const pluginProps = useCloudAppMock
  ? plugin
  : {
      team: process.env.REACT_APP_PLUGIN_TEAM,
      kind: process.env.REACT_APP_PLUGIN_KIND,
      name: process.env.REACT_APP_PLUGIN_NAME,
      version: process.env.REACT_APP_PLUGIN_VERSION,
    };

function App() {
  const { initialValues, initialized, teamName, context } = useFormInit(
    pluginUiMessageHandler,
    false,
  );
  // useFormHeightChange(pluginUiMessageHandler);

  const theme = useMemo(() => createTheme(createThemeOptions()), []);

  const config = useConfig({ initialValues });

  return (
    <PluginContextProvider
      config={config}
      plugin={pluginProps}
      teamName={teamName}
      hideStepper={context === 'wizard'} // TODO: Delete after iframe deprecation
      pluginUiMessageHandler={pluginUiMessageHandler}
      initialValues={initialValues}
    >
      <Box>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <DevWrapper {...devWrapperProps}>
            {initialized && <ConfigUIForm prepareSubmitValues={prepareSubmitValues} />}
          </DevWrapper>
        </ThemeProvider>
      </Box>
    </PluginContextProvider>
  );
}

export default App;
