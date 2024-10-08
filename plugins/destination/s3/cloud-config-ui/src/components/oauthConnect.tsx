import { useState } from 'react';

import { PluginUiMessageHandler } from '@cloudquery/plugin-config-ui-connector';

import {
  AuthType,
  Controller,
  createAndAuthenticateConnector,
  getFieldHelperText,
  useApiCall,
  useFormContext,
  usePluginContext,
} from '@cloudquery/plugin-config-ui-lib';
import CheckIcon from '@mui/icons-material/Check';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import FormControl from '@mui/material/FormControl';
import FormHelperText from '@mui/material/FormHelperText';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

/**
 * @public
 */
export type ConnectProps = {
  pluginUiMessageHandler: PluginUiMessageHandler;
};

/**
 * @public
 * Encapsulatees the Connector logic in a Button or Link.
 */
export function OAuthConnect({ pluginUiMessageHandler }: ConnectProps) {
  const { plugin, teamName } = usePluginContext();
  const { setValue, formState, watch, trigger, getValues } = useFormContext();
  const { callApi } = useApiCall(pluginUiMessageHandler);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const connectorId = watch('connectorId');
  const authType = watch('_authType');
  const [shouldTriggerBucket, setShouldTriggerBucket] = useState(false);

  const handleClick = async () => {
    setShouldTriggerBucket(true);
    await trigger('bucket');

    if (formState.errors.bucket) {
      return;
    }

    setIsLoading(true);

    setValue('connectorId', '');
    setValue('arn', '');
    setValue('externalId', '');

    try {
      const {
        redirect_url: redirectUrl,
        suggested_external_id: suggestedExternalId,
        connectorId: newConnectorId,
      } = await createAndAuthenticateConnector<{
        redirect_url: string;
        suggested_external_id: string;
      }>({
        connectorId,
        pluginName: plugin.name,
        pluginTeamName: plugin.team,
        pluginKind: plugin.kind as any,
        teamName,
        callApi,
        authPluginType: 'aws',
        authenticatePayload: {
          spec: {
            bucket: getValues('bucket'),
          },
        },
      });
      setValue('connectorId', newConnectorId);
      setValue('externalId', suggestedExternalId);

      trigger('connectorId');

      pluginUiMessageHandler.sendMessage('open_url', {
        url: redirectUrl,
      });
    } catch (error: any) {
      setError(error?.body || error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Stack gap={2}>
      <Controller
        name="bucket"
        render={({ field, fieldState }) => (
          <TextField
            disabled={isLoading}
            error={!!fieldState.error}
            fullWidth={true}
            helperText={getFieldHelperText(fieldState.error?.message, 'Name of the S3 bucket.')}
            label="Bucket"
            {...field}
            onChange={(...rest) => {
              field.onChange(...rest);
              if (shouldTriggerBucket) {
                trigger('bucket');
              }

              setValue('connectorId', '');
              setValue('arn', '');
              setValue('externalId', '');
            }}
          />
        )}
      />
      <Stack gap={1}>
        <Box>
          <FormControl>
            <Button
              disabled={!!connectorId || isLoading}
              size="large"
              variant="contained"
              fullWidth={false}
              onClick={handleClick}
              endIcon={connectorId && <CheckIcon />}
            >
              {connectorId
                ? 'AWS Console connected successfully'
                : 'Connect CloudQuery via AWS Console'}
            </Button>
          </FormControl>
        </Box>

        {!error && !formState.errors.connectorId && !isLoading && connectorId && (
          <Typography variant="body2" color="textSecondary">
            To reconnect CloudQuery via AWS IAM Console{' '}
            {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
            <Link underline="always" sx={{ cursor: 'pointer' }} onClick={handleClick}>
              click here
            </Link>
          </Typography>
        )}
        {!error && !formState.errors.connectorId && !isLoading && !connectorId && (
          <Typography variant="body2" color="textSecondary">
            This will open a new browser tab.
          </Typography>
        )}
        {!!error && (
          <FormHelperText error={true} sx={{ marginTop: 2 }}>
            {error.message || 'Something went wrong during authentication. Please try again.'}
          </FormHelperText>
        )}
        {!error && formState.errors.connectorId && (
          <FormHelperText error={true} sx={{ marginTop: 2 }}>
            {`You must authenticate with AWS IAM Console to continue.`}
          </FormHelperText>
        )}
      </Stack>
      <Controller
        name="arn"
        render={({ field, fieldState }) => (
          <TextField
            disabled={isLoading}
            error={!!fieldState.error}
            fullWidth={true}
            helperText={getFieldHelperText(
              fieldState.error?.message,
              authType === AuthType.OAUTH
                ? 'Amazon Resource Name uniquely identifies AWS resources. It will be provided when you finish running the stack'
                : 'Amazon Resource Name uniquely identifies AWS resources. It will be provided when you finish creating the new role',
            )}
            label="ARN"
            {...field}
          />
        )}
      />
      {error && (
        <FormControl>
          {<FormHelperText error={true}>Network error: {error.message}</FormHelperText>}
        </FormControl>
      )}
    </Stack>
  );
}
