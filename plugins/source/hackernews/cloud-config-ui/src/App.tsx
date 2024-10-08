import React from 'react';

import { ConfigUIForm, PluginContextProvider, useFormInit } from '@cloudquery/plugin-config-ui-lib';

import { DevWrapper } from '@cloudquery/plugin-config-ui-lib/components/devWrapper';

import { useConfig } from './hooks/useConfig';
import { envJson } from './utils/envJson';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareSubmitValues } from './utils/prepareSubmitValues';

const getTablesData = () => import('./data/__tables.json');

function App() {
  const { initialValues, initialized, teamName, context } = useFormInit(
    pluginUiMessageHandler,
    true,
  );

  const config = useConfig({ initialValues });

  return (
    <DevWrapper {...envJson}>
      <PluginContextProvider
        config={config}
        teamName={teamName}
        getTablesData={getTablesData}
        hideStepper={context === 'wizard'} // TODO: Delete after iframe deprecation
        pluginUiMessageHandler={pluginUiMessageHandler}
        initialValues={initialValues}
      >
        {initialized && <ConfigUIForm prepareSubmitValues={prepareSubmitValues} />}
      </PluginContextProvider>
    </DevWrapper>
  );
}

export default App;
