import { useCallback, useEffect, useMemo } from 'react';

import { getFieldHelperText, getYupValidationResolver } from '@cloudquery/cloud-ui';
import {
  FormFooter,
  FormWrapper,
  Logo,
  useFormActions,
  useFormCurrentValues,
} from '@cloudquery/plugin-config-ui-lib';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import FormHelperText from '@mui/material/FormHelperText';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import { Controller, FormProvider, Path, useForm } from 'react-hook-form';

import { AdvancedConnectionFields } from './advancedConnectionFields';
import { AdvancedSyncFields } from './advancedSyncFields';
import { FormConnectionFields } from './connectionFields';
import { FormSyncOptions } from './syncOptions';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { parseTestConnectionError } from '../utils/parseTestConnectionError';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues: FormValues | undefined;
  teamName: string;
  plugin: {
    name: string;
    team: string;
    kind: string;
    version: string;
  };
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues, teamName, plugin }: Props) {
  const formContext = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });
  const { control, handleSubmit: handleFormSubmit, getValues, setError, formState } = formContext;

  const getCurrentValues = useCallback(() => prepareSubmitValues(getValues()), [getValues]);
  useFormCurrentValues(pluginUiMessageHandler, getCurrentValues);

  const editMode = !!initialValues?.name;

  const {
    handleCancel,
    handleCancelTestConnection,
    handleDelete,
    handleGoToPreviousStep,
    handleTestConnection,
    handleSubmit,
    isSubmitting,
    isTestingConnection,
    testConnectionError,
    submitPayload,
    submitError,
  } = useFormActions({
    getValues: getCurrentValues,
    teamName,
    pluginUiMessageHandler,
    pluginTeamName: plugin.team,
    pluginName: plugin.name,
    pluginKind: plugin.kind as any,
    pluginVersion: plugin.version,
    isUpdating: editMode,
  });

  useEffect(() => {
    if (submitError) {
      const fieldErrors = submitError.data?.field_errors;

      if (fieldErrors) {
        for (const key of Object.keys(fieldErrors)) {
          if (key in getValues()) {
            setError(key as Path<FormValues>, {
              message: fieldErrors[key],
            });
          } else {
            setError('root', { message: submitError.data.message || submitError.message });

            return;
          }
        }
      } else {
        setError('root', { message: submitError.data.message || submitError.message });
      }
    }
  }, [submitError, getValues, setError]);

  const formDisabled = isSubmitting || isTestingConnection;

  const onTestConnectionSuccess = async () => {
    await handleSubmit(getCurrentValues());
  };

  const onSubmit = handleFormSubmit(handleTestConnection);

  const parsedTestConnectionError = useMemo(
    () => (testConnectionError ? parseTestConnectionError(testConnectionError) : undefined),
    [testConnectionError],
  );

  return (
    <form autoComplete="off" noValidate={true} onSubmit={onSubmit}>
      <Stack spacing={2}>
        <FormWrapper formDisabled={formDisabled}>
          <FormProvider {...formContext}>
            <Stack spacing={2}>
              <Card>
                <CardContent>
                  <Box
                    display="flex"
                    marginBottom={3}
                    justifyContent="space-between"
                    alignItems="center"
                  >
                    <Typography variant="h5">Configure destination</Typography>
                    <Box
                      display="flex"
                      justifyContent="space-between"
                      alignItems="center"
                      gap={1.5}
                    >
                      <Logo src="images/mysql.webp" alt="MySQL" />
                      <Typography variant="body1">MySQL</Typography>
                    </Box>
                  </Box>
                  <Stack marginBottom={2}>
                    <Controller
                      control={control}
                      name="name"
                      render={({ field, fieldState }) => (
                        <TextField
                          error={!!fieldState.error}
                          fullWidth={true}
                          helperText={getFieldHelperText(
                            fieldState.error?.message,
                            'Unique destination name that helps identify the destination within your workspace.',
                          )}
                          label="Destination name"
                          disabled={!!initialValues}
                          autoComplete="off"
                          {...field}
                        />
                      )}
                    />
                  </Stack>
                </CardContent>
              </Card>
              <FormConnectionFields />
              <AdvancedConnectionFields />
              <FormSyncOptions />
              <AdvancedSyncFields />
              <FormHelperText sx={{ textAlign: 'right' }} error={true}>
                {formState.errors.root?.message}
              </FormHelperText>
            </Stack>
          </FormProvider>
        </FormWrapper>
        <FormFooter
          isUpdating={editMode}
          pluginKind={plugin.kind as any}
          isTestingConnection={isTestingConnection}
          isSubmitting={isSubmitting}
          testConnectionError={parsedTestConnectionError}
          submitPayload={submitPayload}
          onCancel={handleCancel}
          onCancelTestConnection={handleCancelTestConnection}
          onTestConnectionSuccess={onTestConnectionSuccess}
          onDelete={handleDelete}
          onGoToPreviousStep={handleGoToPreviousStep}
          submitLabel={`Submit`}
        />
      </Stack>
    </form>
  );
}
