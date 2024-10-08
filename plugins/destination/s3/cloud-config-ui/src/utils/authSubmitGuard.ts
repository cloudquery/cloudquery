import {
  AuthType,
  createAndAuthenticateConnector,
  finishAuthConnectorAuthentication,
  useApiCall,
} from '@cloudquery/plugin-config-ui-lib';

export const authSubmitGuard = async (
  formValues: any,
  teamName: string,
  callApi: ReturnType<typeof useApiCall>['callApi'],
  setValue: (field: string, value: any) => void,
  initialValues: any,
) => {
  try {
    let { connectorId } = formValues;

    if (connectorId === initialValues.connectorId) {
      return true;
    }

    if (formValues._authType === AuthType.OTHER) {
      const { connectorId: newConnectorId } = await createAndAuthenticateConnector<{
        redirect_url: string;
        suggested_external_id: string;
      }>({
        connectorId,
        pluginName: 's3',
        pluginTeamName: 'cloudquery',
        pluginKind: 'destination',
        teamName,
        callApi,
        authPluginType: 'aws',
        authenticatePayload: {
          spec: {
            bucket: formValues.bucket,
          },
        },
      });

      setValue('connectorId', newConnectorId);
      connectorId = newConnectorId;
    }
    await finishAuthConnectorAuthentication({
      connectorId,
      authPluginType: 'aws',
      teamName,
      callApi,
      method: 'PATCH',
      payload: {
        role_arn: formValues.arn,
        external_id: formValues.externalId,
      },
    });

    return true;
  } catch {
    return {
      errorMessage: 'Failed to connect to S3',
    };
  }
};
