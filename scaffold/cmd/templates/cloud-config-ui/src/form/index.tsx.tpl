import { useMemo } from 'react';

import { getFieldHelperText, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import {
  FormFieldGroup,
  generateTablesFromJson,
  Logo,
  useFormSubmit,
} from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

import { Controller, FormProvider, useForm } from 'react-hook-form';

import { PluginTableSelector } from '../components/tableSelector';
import tablesData from '../data/__tables.json';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { prepareInitialValues } from '../utils/prepareInitialValues';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues?: Exclude<FormMessagePayload['init']['initialValues'], undefined>;
}

const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const tablesList = useMemo(() => generateTablesFromJson(tablesData), []);
  const values = useMemo(() => {
    return prepareInitialValues(initialValues, tablesList);
  }, [initialValues, tablesList]);
  const form = useForm<FormValues>({
    defaultValues: values,
    values,
    resolver: formValidationResolver,
  });
  const { control, handleSubmit } = form;

  const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        handleSubmit(resolve, reject)();
      });

      return {
        values: prepareSubmitValues(values, tablesList),
      };
    } catch (error) {
      return { errors: error as Record<string, any> };
    }
  };

  useFormSubmit(handleValidate, pluginUiMessageHandler);

  return (
    <FormProvider {...form}>
      <Stack spacing={2}>
        <Card>
          <CardContent>
            <Stack gap={2}>
              <Box display="flex" justifyContent="space-between" alignItems="center">
                <Typography variant="h5">Configure source</Typography>
                <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
                  <Logo src={`images/icon.png`} alt="{{.Name}}" />
                  <Typography variant="body1">{`{{.Name}}`}</Typography>
                </Box>
              </Box>
              <Stack>
                <Controller
                  control={control}
                  name="name"
                  render={({ field, fieldState }) => (
                    <TextField
                      error={!!fieldState.error}
                      fullWidth={true}
                      helperText={getFieldHelperText(
                        fieldState.error?.message,
                        'Unique source name that helps identify the source within your workspace.',
                      )}
                      label="Source name"
                      disabled={!!initialValues}
                      autoComplete="off"
                      {...field}
                    />
                  )}
                />
              </Stack>
            </Stack>
          </CardContent>
        </Card>
        <FormFieldGroup title="Configuration">
          <Controller
            control={control}
            name="token"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                helperText={fieldState.error?.message}
                label="Token"
                {...field}
              />
            )}
          />
        </FormFieldGroup>
        <FormFieldGroup title="Tables Selection">
          <PluginTableSelector pluginTables={tablesList} />
        </FormFieldGroup>
      </Stack>
    </FormProvider>
  );
}
