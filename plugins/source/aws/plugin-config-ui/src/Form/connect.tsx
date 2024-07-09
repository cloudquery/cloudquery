import { FormFieldGroup } from '@cloudquery/cloud-ui';
import { Box, Button, FormHelperText, Stack, TextField, Typography } from '@mui/material';
import { Controller, useFormContext } from 'react-hook-form';
import { ExclusiveToggle } from './selector';

interface Props {}

export function Connect({}: Props) {
  const form = useFormContext();
  return (
    <FormFieldGroup title="AWS Connection">
      <Stack gap={1}>
        <Box display="flex" justifyContent="space-between" alignItems="center">
          <Typography variant="h5">Connect to AWS</Typography>
          <Box display="flex" justifyContent="space-between" alignItems="center">
            TODO:logo
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
                optionA={{ label: 'AWS Console', value: 'console' }}
                optionB={{ label: 'Manual setup', value: 'manual' }}
                {...field}
              />
            )}
          />
          {form.watch('_setupType') === 'console' && (
            <Stack gap={1}>
              <Box>
                <Button
                  variant="contained"
                  fullWidth={false}
                  onClick={() => {
                    // TODO
                  }}
                >
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
                  fieldState.error?.message ??
                  'It will be provided when you finish running the stack'
                }
                label="ARN"
                {...field}
              />
            )}
          />
          {/* 

      <Controller
        control={control}
        name="spec.username"
        render={({ field, fieldState }) => (
          <Stack direction="row" spacing={2}>
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="Username"
              {...field}
              disabled={defaultUsername === '${username}' && !usernameResetted}
              value={
                defaultUsername === '${username}' && !usernameResetted
                  ? envPlaceholder
                  : field.value
              }
            />
            {defaultUsername === '${username}' && (
              <FormFieldReset
                isResetted={usernameResetted}
                inputSelectorToFocus='input[name="spec.username"]'
                onCancel={() => handelCancelReset('username')}
                onReset={() => handleReset('username')}
              />
            )}
          </Stack>
        )}
      />
      <Controller
        control={control}
        name="spec.password"
        render={({ field, fieldState }) => (
          <Stack direction="row" spacing={2}>
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="Password"
              {...field}
              disabled={defaultPassword === '${password}' && !passwordResetted}
              value={
                defaultPassword === '${password}' && !passwordResetted
                  ? envPlaceholder
                  : field.value
              }
            />
            {defaultPassword === '${password}' && (
              <FormFieldReset
                isResetted={passwordResetted}
                inputSelectorToFocus='input[name="spec.password"]'
                onCancel={() => handelCancelReset('password')}
                onReset={() => handleReset('password')}
              />
            )}
          </Stack>
        )}
      />
      <Controller
        control={control}
        name="spec.clientEncoding"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Client Encoding"
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.ssl"
        render={({ field }) => <FormControlLabel control={<Switch {...field} />} label="SSL" />}
      />
      {sslValue && (
        <Controller
          control={control}
          name="spec.sslMode"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="SSL Mode"
              select={true}
              {...field}
            >
              <MenuItem value={''} hidden={true} />
              {sslModeValues.map((value) => (
                <MenuItem key={value} value={value}>
                  {value}
                </MenuItem>
              ))}
            </TextField>
          )}
        />
      )} */}
        </Stack>
      </Stack>
    </FormFieldGroup>
  );
}
