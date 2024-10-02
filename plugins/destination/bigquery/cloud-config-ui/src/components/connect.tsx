import React from 'react';

import {
  CodeSnippet,
  getFieldHelperText,
  useFormContext,
  usePluginContext,
} from '@cloudquery/plugin-config-ui-lib';
import CheckIcon from '@mui/icons-material/Check';
import { FormControl, FormHelperText } from '@mui/material';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { useGCPConnector } from '../context/GCPConnectorContext';
import { pluginUiMessageHandler } from '../utils/messageHandler';

const redirect = () =>
  pluginUiMessageHandler.sendMessage('open_url', {
    url: 'https://console.cloud.google.com/iam-admin/iam',
  });

export function Connect({ variant = 'button' }: { variant: 'link' | 'button' }) {
  const { plugin, teamName } = usePluginContext();
  const form = useFormContext();
  const connectorId = form.watch('connectorId');
  const serviceAccount = form.watch('_serviceAccount');

  const { createAndAuthenticateConnector, authenticationLoading } = useGCPConnector();

  const getCredentials = async () => {
    const authProps = {
      connectorId,
      pluginName: plugin.name,
      pluginTeamName: plugin.team,
      pluginKind: plugin.kind as any,
      teamName,
    };
    if (connectorId) {
      const { _serviceAccount } = await createAndAuthenticateConnector(authProps);
      form.setValue('_serviceAccount', _serviceAccount);
    } else {
      const { connectorId, _serviceAccount } = await createAndAuthenticateConnector(authProps);

      form.setValue('connectorId', connectorId);
      if (_serviceAccount) {
        form.setValue('_serviceAccount', _serviceAccount);
      }
    }
  };

  const handleClick = async () => {
    if (connectorId) {
      redirect();
    } else {
      await getCredentials();
      redirect();
    }
  };

  return variant === 'button' ? (
    <Stack gap={2}>
      <Stack gap={1}>
        <Box>
          <FormControl>
            <Button
              disabled={!!connectorId || authenticationLoading}
              size="large"
              variant="contained"
              fullWidth={false}
              onClick={handleClick}
              endIcon={connectorId && <CheckIcon />}
            >
              {connectorId ? 'Successfully created GCP credentials' : 'Create credentials'}
            </Button>
            {!!form.formState.errors?.['connectorId']?.message && (
              <FormHelperText error={true}>
                {getFieldHelperText(form.formState.errors?.['connectorId']?.message as string, '')}
              </FormHelperText>
            )}
          </FormControl>
        </Box>

        {connectorId ? (
          <Typography variant="body2" color="textSecondary">
            To reopen the GCP IAM Console {/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
            <Link sx={{ cursor: 'pointer' }} onClick={handleClick}>
              click here
            </Link>
          </Typography>
        ) : (
          <Typography variant="body2" color="textSecondary">
            This will open a new browser tab.
          </Typography>
        )}
      </Stack>

      <Stack>
        <Typography variant="body2">Service Account:</Typography>
        <Box>
          <CodeSnippet text={serviceAccount} />
        </Box>
      </Stack>
    </Stack>
  ) : (
    // eslint-disable-next-line jsx-a11y/anchor-is-valid
    <Link component="button" onClick={handleClick}>
      IAM Console
    </Link>
  );
}
