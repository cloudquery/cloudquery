import {
  Autocomplete,
  Button,
  Checkbox,
  Tab,
  Tabs,
  TextField,
  ToggleButton,
  Typography,
} from '@mui/material';
import { Box, Stack } from '@mui/system';
import { useMemo, useState } from 'react';
import { Controller } from 'react-hook-form';
import { Logo } from '../components/logo';
import { awsServices } from '../utils/constants';

enum ServiceList {
  All = 'all',
  Popular = 'popular',
}

interface Props {}

export function SelectServices({}: Props) {
  const [showServices, setShowServices] = useState<ServiceList.All | ServiceList.Popular>(
    ServiceList.Popular,
  );

  // TODO: where do these come from?
  const TODORegionOpts = ['us-east-1', 'eu-north-1'];
  const TODOServiceOpts = awsServices;

  const filteredServices = useMemo(() => {
    const servicesArray = Object.values(TODOServiceOpts);
    // TODO: filter services by some metric
    return showServices === ServiceList.Popular ? servicesArray.slice(0, 2) : servicesArray;
  }, [TODOServiceOpts, showServices]);

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
          name="regions"
          render={({ field, fieldState }) => {
            // TODO: connect to hook form properly
            console.log({ field });
            return (
              <Autocomplete
                multiple
                id="regions-select"
                options={TODORegionOpts}
                getOptionLabel={(option) => option}
                defaultValue={[]}
                filterSelectedOptions
                renderInput={(params) => (
                  <TextField
                    {...params}
                    {...field}
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
        <Controller
          name="_services"
          render={({ field }) => {
            return (
              <Stack gap={2}>
                <Tabs value={showServices} onChange={(_, newValue) => setShowServices(newValue)}>
                  <Tab label="Popular Services" value={ServiceList.Popular}></Tab>
                  <Tab label="All Services" value={ServiceList.All}></Tab>
                </Tabs>
                <Box
                  display="grid"
                  gap={2}
                  gridTemplateColumns={{ xs: 'minmax(0, 1fr) minmax(0, 1fr)' }}
                  width="100%"
                >
                  {filteredServices.map((service) => {
                    const isChecked = field.value?.includes(service.name);

                    return (
                      <ToggleButton
                        sx={{ display: 'flex', justifyContent: 'space-between', py: 0.5, pr: 0 }}
                        key={service.name}
                        value={service.name}
                        onClick={() =>
                          field.onChange(() =>
                            isChecked
                              ? field.value.filter((name: string) => name !== service.name)
                              : [...field.value, service.name],
                          )
                        }
                      >
                        <Box display="flex" alignItems="center" gap={1}>
                          <Logo
                            src={service.logo}
                            fallbackSrc="/images/aws.webp"
                            alt={service.name}
                            height={32}
                            width={32}
                          />
                          <Typography variant="body1">{service.name}</Typography>
                        </Box>
                        <Checkbox checked={isChecked} />
                      </ToggleButton>
                    );
                  })}
                </Box>
                <Button
                  fullWidth
                  onClick={() =>
                    setShowServices(
                      showServices === ServiceList.Popular ? ServiceList.All : ServiceList.Popular,
                    )
                  }
                >
                  {showServices === ServiceList.Popular
                    ? 'Show all services'
                    : 'Show only popular services'}
                </Button>
              </Stack>
            );
          }}
        />
      </Stack>
    </Stack>
  );
}
