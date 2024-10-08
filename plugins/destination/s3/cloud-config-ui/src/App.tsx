import { useEffect, useState } from 'react';

import {
  ConfigUIForm,
  getAuthenticateConnector,
  PluginContextProvider,
  useApiCall,
  useFormInit,
} from '@cloudquery/plugin-config-ui-lib';
import { DevWrapper } from '@cloudquery/plugin-config-ui-lib/components/devWrapper';

import { useConfig } from './hooks/useConfig';
import { envJson } from './utils/envJson';
import { pluginUiMessageHandler } from './utils/messageHandler';
import { prepareSubmitValues } from './utils/prepareSubmitValues';

function App() {
  const { initialValues, initialized, teamName, context } = useFormInit(
    pluginUiMessageHandler,
    true,
  );
  const [isLoading, setIsLoading] = useState(true);
  const [authenticateConnector, setAuthenticateConnector] = useState<
    | {
        arn: string;
        externalId: string;
      }
    | undefined
  >(undefined);
  const { callApi } = useApiCall(pluginUiMessageHandler);

  useEffect(() => {
    if (initialValues?.connectorId) {
      getAuthenticateConnector({
        callApi,
        connectorId: initialValues.connectorId,
        teamName,
        authPluginType: 'aws',
      })
        .then((response) => {
          setAuthenticateConnector({
            arn: response.body.role_arn,
            externalId: response.body.external_id,
          });
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else {
      setIsLoading(false);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [initialValues]);

  const parsedInitialValues = initialValues
    ? {
        ...initialValues,
        spec: {
          ...initialValues.spec,
          ...authenticateConnector,
        },
      }
    : undefined;

  const config = useConfig({
    initialValues: parsedInitialValues,
  });

  return (
    <DevWrapper {...envJson}>
      <PluginContextProvider
        config={config}
        teamName={teamName}
        hideStepper={context === 'wizard'} // TODO: Delete after iframe deprecation
        pluginUiMessageHandler={pluginUiMessageHandler}
        initialValues={parsedInitialValues}
      >
        {initialized && !isLoading && <ConfigUIForm prepareSubmitValues={prepareSubmitValues} />}
      </PluginContextProvider>
    </DevWrapper>
  );
}

export default App;
