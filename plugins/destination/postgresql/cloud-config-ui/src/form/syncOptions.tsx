import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';
import Stack from '@mui/material/Stack';
import { Controller, useFormContext } from 'react-hook-form';
import { FormValues, migrateModeValues, writeModeValues } from '../utils/formSchema';
import { FormFieldGroup } from '@cloudquery/plugin-config-ui-lib';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import Link from '@mui/material/Link';
import { getFieldHelperText } from '@cloudquery/cloud-ui';

export function FormSyncOptions() {
  const { control } = useFormContext<FormValues>();

  return (
    <FormFieldGroup
      title="Sync Options"
      subheader="Configure how CloudQuery should write to your destination."
    >
      <Stack spacing={3}>
        <Controller
          control={control}
          name="migrateMode"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              required={true}
              helperText={getFieldHelperText(
                fieldState.error?.message,
                <>
                  Specifies the migration mode to use when source tables are changed.{' '}
                  <Link
                    href="#"
                    onClick={(event) => {
                      event.preventDefault();
                      pluginUiMessageHandler.sendMessage('open_url', {
                        url: 'https://docs.cloudquery.io/docs/reference/destination-spec#migrate_mode',
                      });
                    }}
                  >
                    Read more
                  </Link>
                </>,
              )}
              label="Migrate mode"
              select={true}
              SelectProps={{
                MenuProps: {
                  autoFocus: false,
                  disableAutoFocus: true,
                },
              }}
              {...field}
            >
              {migrateModeValues.map((value) => (
                <MenuItem key={value} value={value}>
                  {value}
                </MenuItem>
              ))}
            </TextField>
          )}
        />
        <Controller
          control={control}
          name="writeMode"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              required={true}
              helperText={getFieldHelperText(
                fieldState.error?.message,
                <>
                  Specifies the update method to use when inserting rows.{' '}
                  <Link
                    href="#"
                    onClick={(event) => {
                      event.preventDefault();
                      pluginUiMessageHandler.sendMessage('open_url', {
                        url: 'https://docs.cloudquery.io/docs/reference/destination-spec#write_mode',
                      });
                    }}
                  >
                    Read more
                  </Link>
                </>,
              )}
              label="Write mode"
              select={true}
              SelectProps={{
                MenuProps: {
                  autoFocus: false,
                  disableAutoFocus: true,
                },
              }}
              {...field}
            >
              {writeModeValues.map((value) => (
                <MenuItem key={value} value={value}>
                  {value}
                </MenuItem>
              ))}
            </TextField>
          )}
        />
      </Stack>
    </FormFieldGroup>
  );
}
