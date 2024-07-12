import { Box, Button, Grid, Stack } from '@mui/material';
import { useForm, FormProvider, SubmitHandler } from 'react-hook-form';
import { FormFieldGroup, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useState } from 'react';

import { Connect } from './connect';
import { AWSFormStepper } from '../components/stepper';
import { Guides } from '../components/guides';
import { SelectServices } from './selectServices';
import { useGetAWSServices } from '../hooks/useGetAWSServices';
import { useAuthenticateConnectorFinishAWS } from '../hooks/useAuthenticateAWSFinish';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const [activeIndex, setActiveIndex] = useState(0);
  const form = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });
  const {
    control,
    handleSubmit,
    formState: { defaultValues },
    setValue,
    watch,
  } = form;
  const awsServices = useGetAWSServices();

  const { mutateAsync: finishAWSAuth } = useAuthenticateConnectorFinishAWS();

  // const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
  //   try {
  //     const values: FormValues = await new Promise((resolve, reject) => {
  //       // TODO: this should trigger AuthenticateConnectorFinishAWS prior to Testing Connection
  //       handleSubmit(resolve, reject)();
  //     });

  //     return {
  //       values: prepareSubmitValues(values, awsServices),
  //     };
  //   } catch (error) {
  //     return { errors: error as Record<string, any> };
  //   }
  // };

  // useFormSubmit(handleValidate, pluginUiMessageHandler);

  const onSubmit: SubmitHandler<FormValues> = async (values) => {
    try {
      finishAWSAuth({
        connectorId: values.connector_id,
        data: { role_arn: values.arn },
      });

      // then actual submission
      //prepareSubmitValues(values, awsServices),
    } catch (e) {}
  };
  const handleCancel = () => {};

  return (
    <form autoComplete="off" noValidate={true} onSubmit={handleSubmit(onSubmit)}>
      <Stack gap={5}>
        <FormProvider {...form}>
          <AWSFormStepper activeIndex={activeIndex} setActiveIndex={setActiveIndex} />
          <Grid container spacing={2}>
            <Grid item xs={7} md={6}>
              <FormFieldGroup title="AWS Connection">
                <Stack gap={2}>
                  <Box sx={{ display: activeIndex === 0 ? 'block' : 'none' }}>
                    <Connect />
                  </Box>
                  <Box sx={{ display: activeIndex === 1 ? 'block' : 'none' }}>
                    <SelectServices awsServices={awsServices} />
                  </Box>
                  <Box display="flex" justifyContent="space-between">
                    <Box>
                      {activeIndex > 0 && (
                        <Button
                          onClick={() => setActiveIndex(activeIndex - 1)}
                          variant="text"
                          color="secondary"
                        >
                          Previous step
                        </Button>
                      )}
                    </Box>
                    <Box display="flex" gap={2}>
                      <Button variant="text" onClick={handleCancel}>
                        Cancel
                      </Button>
                      <Button type="submit" variant="contained">
                        Submit:TODO
                      </Button>
                    </Box>
                  </Box>
                </Stack>
              </FormFieldGroup>
            </Grid>
            <Grid item xs={5} md={6}>
              <Guides />
            </Grid>
          </Grid>
        </FormProvider>
      </Stack>
    </form>
  );
}
