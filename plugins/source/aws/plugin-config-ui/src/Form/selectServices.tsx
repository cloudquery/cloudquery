import { FormFieldGroup } from '@cloudquery/cloud-ui';
import {
  Autocomplete,
  Button,
  Checkbox,
  Grid,
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

interface Props {}

export function SelectServices({}: Props) {
  const [showServices, setShowServices] = useState<'popular' | 'all'>('popular');

  const TODORegionOpts = ['us-east-1', 'eu-north-1'];
  const TODOServiceOpts = ['EC2', 'RDS', 'S3', 'DynamoDB'];
  const filteredServices = useMemo(() => {
    return showServices === 'popular' ? TODOServiceOpts.slice(0, 2) : TODOServiceOpts;
  }, [TODOServiceOpts, showServices]);

  return (
    <FormFieldGroup title="AWS Connection">
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
            name="services"
            render={({ field, fieldState }) => {
              // TODO: connect to hook form properly
              console.log({ field, fieldState });
              return (
                <Stack gap={2}>
                  <Tabs value={showServices} onChange={(_, newValue) => setShowServices(newValue)}>
                    <Tab label="Popular Services" value="popular"></Tab>
                    <Tab label="All Services" value="all"></Tab>
                  </Tabs>
                  <Box display="grid" gap={2} gridTemplateColumns={{ xs: '1fr 1fr' }} width="100%">
                    {filteredServices.map((service) => (
                      <ToggleButton
                        sx={{ display: 'flex', justifyContent: 'space-between', py: 0.5, pr: 0 }}
                        key={service}
                        value={service}
                      >
                        <Box display="flex" alignItems="center" gap={1}>
                          <Logo src="/images/aws.webp" alt={service} height={32} width={32} />
                          <Typography variant="body1">{service}</Typography>
                        </Box>
                        <Checkbox checked={field.value?.includes(service)} />
                      </ToggleButton>
                    ))}
                  </Box>
                  <Button
                    fullWidth
                    onClick={() => setShowServices(showServices === 'popular' ? 'all' : 'popular')}
                  >
                    {showServices === 'popular'
                      ? 'Show all services'
                      : 'Show only popular services'}
                  </Button>
                </Stack>
              );
            }}
          />
        </Stack>
      </Stack>
    </FormFieldGroup>

    // <FormFieldGroup title="PostgreSQL Connection">
    //   <Controller
    //     control={control}
    //     name="spec.host"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Host"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.port"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Port"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.database"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Database"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.username"
    //     render={({ field, fieldState }) => (
    //       <Stack direction="row" spacing={2}>
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="Username"
    //           {...field}
    //           disabled={defaultUsername === '${username}' && !usernameResetted}
    //           value={
    //             defaultUsername === '${username}' && !usernameResetted
    //               ? envPlaceholder
    //               : field.value
    //           }
    //         />
    //         {defaultUsername === '${username}' && (
    //           <FormFieldReset
    //             isResetted={usernameResetted}
    //             inputSelectorToFocus='input[name="spec.username"]'
    //             onCancel={() => handelCancelReset('username')}
    //             onReset={() => handleReset('username')}
    //           />
    //         )}
    //       </Stack>
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.password"
    //     render={({ field, fieldState }) => (
    //       <Stack direction="row" spacing={2}>
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="Password"
    //           {...field}
    //           disabled={defaultPassword === '${password}' && !passwordResetted}
    //           value={
    //             defaultPassword === '${password}' && !passwordResetted
    //               ? envPlaceholder
    //               : field.value
    //           }
    //         />
    //         {defaultPassword === '${password}' && (
    //           <FormFieldReset
    //             isResetted={passwordResetted}
    //             inputSelectorToFocus='input[name="spec.password"]'
    //             onCancel={() => handelCancelReset('password')}
    //             onReset={() => handleReset('password')}
    //           />
    //         )}
    //       </Stack>
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.clientEncoding"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Client Encoding"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.ssl"
    //     render={({ field }) => <FormControlLabel control={<Switch {...field} />} label="SSL" />}
    //   />
    //   {sslValue && (
    //     <Controller
    //       control={control}
    //       name="spec.sslMode"
    //       render={({ field, fieldState }) => (
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="SSL Mode"
    //           select={true}
    //           {...field}
    //         >
    //           <MenuItem value={''} hidden={true} />
    //           {sslModeValues.map((value) => (
    //             <MenuItem key={value} value={value}>
    //               {value}
    //             </MenuItem>
    //           ))}
    //         </TextField>
    //       )}
    //     />
    //   )}
    // </FormFieldGroup>
  );
}
