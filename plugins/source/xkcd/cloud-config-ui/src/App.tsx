import { Fragment, useEffect, useState } from 'react';

import {
  CloudAppMock,
  CloudQueryTables,
  ConfigUIForm,
  PluginContextProvider,
  useFormInit,
} from '@cloudquery/plugin-config-ui-lib';

import { useConfig } from './hooks/useConfig';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareSubmitValues } from './utils/prepareSubmitValues';

const useCloudAppMock =
  (process.env.REACT_APP_USE_CLOUD_APP_MOCK === 'true' || process.env.NODE_ENV !== 'production') &&
  globalThis.self === globalThis.top;
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
  const [tablesData, setTablesData] = useState<CloudQueryTables | undefined>();
  useEffect(() => {
    import('./data/__tables.json').then(({ default: data }) => setTablesData(data));
  }, []);

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
      tablesData={tablesData as CloudQueryTables}
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
