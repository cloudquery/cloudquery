import { Grid, Stack } from '@mui/material';
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
import { useGetPlugin } from '../hooks/useGetPlugin';

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

  const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        handleSubmit(resolve, reject)();
      });

      return {
        values: prepareSubmitValues(values),
      };
    } catch (error) {
      return { errors: error as Record<string, any> };
    }
  };

  useFormSubmit(handleValidate, pluginUiMessageHandler);

  const { data: pluginData, error: pluginDataError, isLoading: pluginDataLoading } = useGetPlugin();
  console.log({ pluginData });

  return (
    <Stack gap={5}>
      <FormProvider {...form}>
        <AWSFormStepper activeIndex={activeIndex} setActiveIndex={setActiveIndex} />
        <Grid container spacing={2}>
          <Grid item xs={7} md={6}>
            <FormFieldGroup title="AWS Connection">
              {activeIndex === 0 && <Connect />}
              {activeIndex === 1 && <SelectServices />}
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
