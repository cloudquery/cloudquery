import { Autocomplete, TextField, Typography } from '@mui/material';
import { Box, Stack } from '@mui/system';
import { Controller } from 'react-hook-form';
import { Logo } from '../components/logo';
import { awsRegions, top8Services } from '../utils/constants';
import { ServiceList, ServiceType } from '../components/serviceList';

interface Props {
  awsServices: Record<string, ServiceType>;
}

export function SelectServices({ awsServices: serviceOptions }: Props) {
  const regionOptions = awsRegions;

  return (
    <Stack gap={1}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h5">Select regions and services</Typography>
        <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
          <Logo src="/images/aws.webp" alt="AWS" />
          <Typography variant="body1">AWS</Typography>
        </Box>
      </Box>
      <Typography variant="caption" color="secondary">
        Select services you want to sync your data from
      </Typography>
      <Stack gap={3}>
        <Controller
          name="spec.regions"
          render={({ field, fieldState }) => {
            const { onChange, value, ...fieldProps } = field;
            return (
              <Autocomplete
                multiple
                id="regions-select"
                options={regionOptions}
                getOptionLabel={(option) => option}
                value={value}
                onChange={(_, newValue) => {
                  onChange(newValue);
                }}
                filterSelectedOptions
                renderInput={(params) => (
                  <TextField
                    {...params}
                    {...fieldProps}
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={fieldState.error?.message}
                    label="Regions"
                  />
                )}
              />
            );
          }}
        />
        <ServiceList
          topServices={top8Services}
          services={serviceOptions}
          formControlName="services"
          fallbackLogoSrc="/images/aws.webp"
        />
      </Stack>
    </Stack>
  );
}
