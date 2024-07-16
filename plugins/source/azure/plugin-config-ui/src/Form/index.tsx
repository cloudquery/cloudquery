import { Box, Grid, Stack } from '@mui/material';
import { useForm, FormProvider, SubmitHandler } from 'react-hook-form';
import { FormFieldGroup, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';

import { Connect } from './connect';
import { AzureFormStepper } from '../components/stepper';
import { Guides } from '../components/guides';
import { SelectServices } from './selectServices';
import { useGetServices } from '../hooks/useGetServices';
import { Footer } from './footer';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
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
  const activeIndex = watch('_activeIndex');
  const services = useGetServices();

  const onSubmit: SubmitHandler<FormValues> = async (values) => {
    try {
      prepareSubmitValues(values, services);
      // TODO:SUBMIT
    } catch (e) {
      // TODO:ERROR
    }
  };
  const handleCancel = () => {
    // TODO:FOOTER
  };
  const handlePreviousStep = () => {
    // TODO:FOOTER
  };

  return (
    <form autoComplete="off" noValidate={true} onSubmit={handleSubmit(onSubmit)}>
      <Stack gap={5}>
        <FormProvider {...form}>
          <AzureFormStepper />
          <Grid container spacing={2}>
            <Grid item xs={7} md={6}>
              <FormFieldGroup>
                <Stack gap={2}>
                  <Box sx={{ display: activeIndex === 0 ? 'block' : 'none' }}>
                    <Connect />
                  </Box>
                  <Box sx={{ display: activeIndex === 1 ? 'block' : 'none' }}>
                    <SelectServices services={services} />
                  </Box>
                  <Footer handleCancel={handleCancel} handlePreviousStep={handlePreviousStep} />
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
