import { Box, Button, FormHelperText, Stack, TextField, Typography } from '@mui/material';
import { Controller, useFormContext } from 'react-hook-form';
import { ExclusiveToggle } from '../components/selector';
import { Logo } from '../components/logo';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { SetupType } from '../utils/formSchema';
import { useAuthenticateConnectorAWS } from '../hooks/useAuthenticateAWS';

interface Props {}

export function Connect({}: Props) {
  const form = useFormContext();

  const { mutateAsync } = useAuthenticateConnectorAWS({});

  const handleClick = async () => {
    const rsp = await mutateAsync({
      connectorId: 'abc',
      data: {
        plugin_kind: 'source',
        plugin_name: 'aws',
        plugin_team: 'cloudquery',
      },
    });
    console.log({ rsp });

    pluginUiMessageHandler.sendMessage('open_url', {
      url: rsp.redirect_url,
    });
  };

  return (
    <Stack gap={1}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h5">Connect to AWS</Typography>
        <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
          <Logo src="/images/aws.webp" alt="AWS" />
          <Typography variant="body1">AWS</Typography>
        </Box>
      </Box>
      <Typography variant="body2">
        To securely connect to AWS, we require a Cross-Account IAM Role to be created:
      </Typography>
      <Stack gap={3}>
        <Controller
          name="name"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="Source name"
              {...field}
            />
          )}
        />
        <Controller
          name="_setupType"
          render={({ field }) => (
            <ExclusiveToggle
              optionA={{ label: 'AWS Console', value: SetupType.Console }}
              optionB={{ label: 'Manual setup', value: SetupType.Manual }}
              {...field}
            />
          )}
        />
        {form.watch('_setupType') === SetupType.Console && (
          <Stack gap={1}>
            <Box>
              <Button variant="contained" fullWidth={false} onClick={handleClick}>
                Connect CloudQuery via AWS Console
              </Button>
            </Box>

            <FormHelperText>This will open a new browser tab.</FormHelperText>
          </Stack>
        )}
        <Controller
          name="arn"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={
                fieldState.error?.message ?? 'It will be provided when you finish running the stack'
              }
              label="ARN"
              {...field}
            />
          )}
        />
      </Stack>
    </Stack>
  );
}
