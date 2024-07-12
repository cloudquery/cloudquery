import { Button, Grid, Stack } from '@mui/material';
import { useForm, FormProvider } from 'react-hook-form';
import { FormFieldGroup, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useState } from 'react';
import { useFormSubmit } from '@cloudquery/plugin-config-ui-lib';

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

  const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        // TODO: this should trigger AuthenticateConnectorFinishAWS prior to Testing Connection
        handleSubmit(resolve, reject)();
      });

      return {
        values: prepareSubmitValues(values, awsServices),
      };
    } catch (error) {
      return { errors: error as Record<string, any> };
    }
  };

  // TODO: move this into Test Connection as pre-req
  const doThing = () => {
    finishAWSAuth({
      connectorId: form.getValues('connector_id'),
      data: { role_arn: form.getValues('arn') },
    });
  };

  useFormSubmit(handleValidate, pluginUiMessageHandler);

  return (
    <Stack gap={5}>
      <FormProvider {...form}>
        <AWSFormStepper activeIndex={activeIndex} setActiveIndex={setActiveIndex} />
        <Grid container spacing={2}>
          <Grid item xs={7} md={6}>
            <FormFieldGroup title="AWS Connection">
              {activeIndex === 0 && <Connect />}
              {activeIndex === 1 && <SelectServices awsServices={awsServices} />}
              <Button onClick={doThing}>TODO:Do Thing</Button>
            </FormFieldGroup>
          </Grid>
          <Grid item xs={5} md={6}>
            <Guides />
          </Grid>
        </Grid>
      </FormProvider>
    </Stack>
  );
}
