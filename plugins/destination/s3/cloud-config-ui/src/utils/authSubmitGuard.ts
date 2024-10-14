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
) => {
  try {
    const {
      connectorId,
      externalId,
      arn,
      _finishedConnectorId,
      _finishedExternalId,
      _finishedArn,
    } = formValues;
    let connectorIdValue = connectorId;

    const finishedValuesMatch =
      _finishedConnectorId &&
      _finishedExternalId &&
      _finishedArn &&
      _finishedConnectorId === connectorId &&
      _finishedExternalId === externalId &&
      _finishedArn === arn;

    if (finishedValuesMatch) {
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

      connectorIdValue = newConnectorId;
      setValue('connectorId', connectorIdValue);
    }
    await finishAuthConnectorAuthentication({
      connectorId: connectorIdValue,
      authPluginType: 'aws',
      teamName,
      callApi,
      method: 'PATCH',
      payload: {
        role_arn: formValues.arn,
        external_id: formValues.externalId,
      },
    });

    setValue('_finishedConnectorId', connectorIdValue);
    setValue('_finishedExternalId', formValues.externalId);
    setValue('_finishedArn', formValues.arn);

    return true;
  } catch {
    return {
      errorMessage: 'Failed to connect to S3',
    };
  }
};
