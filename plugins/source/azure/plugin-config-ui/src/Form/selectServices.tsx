import { Typography } from '@mui/material';
import { Box, Stack } from '@mui/system';
import { Logo } from '../components/todoShare/logo';
import { top8Services } from '../utils/constants';
import { ServiceList, ServiceTypes } from '../components/todoShare/serviceList';

interface Props {
  services: ServiceTypes;
}

const fallbackLogoSrc = '/images/azure.webp';

export function SelectServices({ services: serviceOptions }: Props) {
  return (
    <Stack gap={1}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h5">Select services</Typography>
        <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
          <Logo src={fallbackLogoSrc} alt="Azure" />
          <Typography variant="body1">Azure</Typography>
        </Box>
      </Box>
      <Typography variant="caption" color="secondary">
        Select services you want to sync your data from
      </Typography>
      <Stack gap={3}>
        <ServiceList
          topServices={top8Services}
          services={serviceOptions}
          formControlName="services"
          fallbackLogoSrc={fallbackLogoSrc}
        />
      </Stack>
    </Stack>
  );
}
