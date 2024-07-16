import {
  Box,
  Button,
  Checkbox,
  Stack,
  Tab,
  Tabs,
  ToggleButton,
  Tooltip,
  Typography,
  useTheme,
} from '@mui/material';
import { useEffect, useMemo, useState } from 'react';
import { Logo } from './logo';
import { Controller } from 'react-hook-form';

enum ServiceListMode {
  All = 'all',
  Popular = 'popular',
}

export type ServiceType = {
  name: string;
  label: string;
  logo: string;
  tables: string[];
};

interface Props {
  services: Record<string, ServiceType>;
  topServices: string[];
  formControlName: string;
  fallbackLogoSrc?: string;
}

export function ServiceList({ services, topServices, formControlName, fallbackLogoSrc }: Props) {
  const { palette } = useTheme();

  const [showServices, setShowServices] = useState<ServiceListMode.All | ServiceListMode.Popular>(
    ServiceListMode.Popular,
  );

  // prefetch logos
  useEffect(() => {
    if (services) {
      Object.values(services).forEach((service) => {
        const img = new Image();
        img.src = service.logo;
      });
    }
  }, [services]);

  const filteredServices: ServiceType[] = useMemo(
    () =>
      showServices === ServiceListMode.Popular
        ? topServices.map((name) => services[name]).filter(Boolean)
        : Object.values(services).sort((a, b) => a.label.localeCompare(b.label)),
    [services, showServices, topServices],
  );

  return (
    <Controller
      name={formControlName}
      render={({ field }) => {
        return (
          <Stack gap={2}>
            <Tabs value={showServices} onChange={(_, newValue) => setShowServices(newValue)}>
              <Tab label="Popular Services" value={ServiceListMode.Popular}></Tab>
              <Tab label="All Services" value={ServiceListMode.All}></Tab>
            </Tabs>
            <Box
              display="grid"
              gap={2}
              gridTemplateColumns={{ xs: 'minmax(0, 1fr) minmax(0, 1fr)' }}
              width="100%"
              maxHeight="calc(100vh - 550px)"
              sx={{ overflowY: 'auto' }}
            >
              {filteredServices.map((service) => {
                const isChecked = field.value?.includes(service.name);

                return (
                  <ToggleButton
                    sx={{
                      display: 'flex',
                      justifyContent: 'space-between',
                      py: 0.5,
                      pr: 0,
                    }}
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
                    <Box
                      display="flex"
                      alignItems="center"
                      gap={1}
                      justifyContent="space-between"
                      width="100%"
                    >
                      <Box display="flex" alignItems="center" gap={1} flexShrink={1} width="70%">
                        <Logo
                          src={service.logo}
                          fallbackSrc={fallbackLogoSrc}
                          alt={service.name}
                          height={32}
                          width={32}
                        />
                        <Tooltip title={service.label}>
                          <Typography
                            sx={{
                              overflow: 'hidden',
                              textOverflow: 'ellipsis',
                              whiteSpace: 'nowrap',
                            }}
                            color={palette.grey[400]}
                            fontWeight="bold"
                            variant="body1"
                          >
                            {service.label}
                          </Typography>
                        </Tooltip>
                      </Box>
                      <Checkbox checked={isChecked} />
                    </Box>
                  </ToggleButton>
                );
              })}
            </Box>
            <Button
              fullWidth
              onClick={() =>
                setShowServices(
                  showServices === ServiceListMode.Popular
                    ? ServiceListMode.All
                    : ServiceListMode.Popular,
                )
              }
            >
              {showServices === ServiceListMode.Popular
                ? 'Show all services'
                : 'Show only popular services'}
            </Button>
          </Stack>
        );
      }}
    />
  );
}
