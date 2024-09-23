import React, { Fragment, Suspense } from 'react';

import { ConfigUIForm, PluginContextProvider, useFormInit } from '@cloudquery/plugin-config-ui-lib';

import { useConfig } from './hooks/useConfig';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareSubmitValues } from './utils/prepareSubmitValues';

const CloudAppMock: React.FC<any> = React.lazy(() =>
  import('@cloudquery/plugin-config-ui-lib/dist/components/cloudAppMock').then(
    ({ CloudAppMock }) => ({
      default: CloudAppMock,
    }),
  ),
);

const CloudAppMockWrapper = (props: any) => (
  <Suspense>
    <CloudAppMock {...props} />
  </Suspense>
);

const useCloudAppMock =
  (process.env.REACT_APP_USE_CLOUD_APP_MOCK === 'true' || process.env.NODE_ENV !== 'production') &&
  window.self === window.top;
const DevWrapper = useCloudAppMock ? CloudAppMockWrapper : Fragment;
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
    true,
  );

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
      <DevWrapper {...devWrapperProps}>
        {initialized && <ConfigUIForm prepareSubmitValues={prepareSubmitValues} />}
      </DevWrapper>
    </PluginContextProvider>
  );
}

export default App;
