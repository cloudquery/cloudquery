import { Box, Stack, TextField, Typography } from '@mui/material';
import { Controller, useFormContext } from 'react-hook-form';
import { Logo } from '../components/todoShare/logo';
import { pluginUiMessageHandler } from '../utils/messageHandler';

interface Props {}

export function Connect({}: Props) {
  const form = useFormContext();

  const handleClick = async () => {
    // TODO:SUBMIT
  };

  return (
    <Stack gap={3}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h5">Connect to Azure</Typography>
        <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
          <Logo src="/images/azure.webp" alt="Azure" />
          <Typography variant="body1">Azure</Typography>
        </Box>
      </Box>

      <Stack gap={2}>
        <Controller
          name="name"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="Source name"
              {...field}
              sx={{ mb: 1 }}
            />
          )}
        />

        <Controller
          name="env.AZURE_TENANT_ID"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={
                fieldState.error?.message ?? (
                  <>
                    This is the value of <b>tenant</b> in the JSON output
                  </>
                )
              }
              label="Azure AD Tenant ID"
              {...field}
            />
          )}
        />
        <Controller
          name="env.AZURE_CLIENT_ID"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={
                fieldState.error?.message ?? (
                  <>
                    This is the value of <b>appId</b> in the JSON output
                  </>
                )
              }
              label="Service Principal App ID"
              {...field}
            />
          )}
        />
        <Controller
          name="env.AZURE_CLIENT_SECRET"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={
                fieldState.error?.message ?? (
                  <>
                    This is the value of <b>password</b> in the JSON output
                  </>
                )
              }
              label="Service Principal Password"
              {...field}
            />
          )}
        />
      </Stack>
    </Stack>
  );
}
