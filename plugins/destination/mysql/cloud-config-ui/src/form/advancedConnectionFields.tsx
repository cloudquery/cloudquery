import { getFieldHelperText } from '@cloudquery/cloud-ui';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

import Accordion from '@mui/material/Accordion';
import AccordionDetails from '@mui/material/AccordionDetails';
import AccordionSummary from '@mui/material/AccordionSummary';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import FormControl from '@mui/material/FormControl';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormHelperText from '@mui/material/FormHelperText';
import MenuItem from '@mui/material/MenuItem';
import Stack from '@mui/material/Stack';
import Switch from '@mui/material/Switch';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import { Controller, useFormContext } from 'react-hook-form';

import { FormValues, tlsModeValues } from '../utils/formSchema';

export function AdvancedConnectionFields() {
  const { control, watch } = useFormContext<FormValues>();

  const tlsEnabled = watch('connectionParams.tls');
  const parseTimeEnabled = watch('connectionParams.parseTime');
  const connectionType = watch('connectionType');

  return connectionType === 'fields' ? (
    <Card>
      <CardContent>
        <Accordion
          disableGutters={true}
          sx={{
            '&:before': {
              display: 'none',
            },
            boxShadow: 'none',
            backgroundColor: 'transparent',
          }}
        >
          <AccordionSummary
            sx={{ backgroundColor: 'transparent', paddingLeft: 0 }}
            expandIcon={<ExpandMoreIcon />}
          >
            <Typography variant="h6">Advanced Connection Options</Typography>
          </AccordionSummary>
          <AccordionDetails sx={{ paddingLeft: 0 }}>
            <Stack spacing={2}>
              <Controller
                control={control}
                name="tcp"
                render={({ field, fieldState }) => (
                  <FormControl>
                    <FormControlLabel
                      control={<Switch checked={field.value} {...field} />}
                      label="TCP"
                    />
                    <FormHelperText error={!!fieldState.error?.message}>
                      {getFieldHelperText(
                        fieldState.error?.message,
                        'If true, will enable connection over TCP to the server. Optional, defaults to false.',
                      )}
                    </FormHelperText>
                  </FormControl>
                )}
              />
              <Controller
                control={control}
                name="connectionParams.tls"
                render={({ field, fieldState }) => (
                  <FormControl>
                    <FormControlLabel
                      control={<Switch checked={field.value} {...field} />}
                      label="TLS"
                    />
                    <FormHelperText error={!!fieldState.error?.message}>
                      {getFieldHelperText(
                        fieldState.error?.message,
                        'If true, will enabled TLS/SSL encrypted connection to the server. Optional, defaults to false.',
                      )}
                    </FormHelperText>
                  </FormControl>
                )}
              />
              {tlsEnabled && (
                <Controller
                  control={control}
                  name="connectionParams.tlsMode"
                  render={({ field, fieldState }) => (
                    <TextField
                      error={!!fieldState.error}
                      fullWidth={true}
                      helperText={getFieldHelperText(
                        fieldState.error?.message,
                        'SSL connections to encrypt client/server communications using TLS protocols for increased security.',
                      )}
                      label="TLS Mode"
                      select={true}
                      SelectProps={{
                        MenuProps: {
                          autoFocus: false,
                          disableAutoFocus: true,
                        },
                      }}
                      {...field}
                    >
                      <MenuItem value={''} hidden={true} />
                      {tlsModeValues.map((value) => (
                        <MenuItem key={value} value={value}>
                          {value}
                        </MenuItem>
                      ))}
                    </TextField>
                  )}
                />
              )}
              <Controller
                control={control}
                name="connectionParams.parseTime"
                render={({ field, fieldState }) => (
                  <FormControl>
                    <FormControlLabel
                      control={<Switch checked={field.value} {...field} />}
                      label="Parse Time"
                    />
                    <FormHelperText error={!!fieldState.error?.message}>
                      {getFieldHelperText(
                        fieldState.error?.message,
                        'If true, changes the output type of DATE and DATETIME values to time.Time instead of []byte / string. Optional, defaults to false.',
                      )}
                    </FormHelperText>
                  </FormControl>
                )}
              />
              {parseTimeEnabled && (
                <Controller
                  control={control}
                  name="connectionParams.loc"
                  render={({ field, fieldState }) => (
                    <TextField
                      error={!!fieldState.error}
                      fullWidth={true}
                      helperText={getFieldHelperText(
                        fieldState.error?.message,
                        `Sets the location for time.Time values. "Local" sets the system's location. Optional, defaults to UTC.`,
                      )}
                      label="Location"
                      autoComplete="off"
                      {...field}
                    />
                  )}
                />
              )}

              <Controller
                control={control}
                name="connectionParams.charset"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Sets the charset used for client-server interaction. Multiple charsets can be configured with comma separation (ex. utf8mb4,utf8). Optional, defaults to utf8mb4.',
                    )}
                    label="Charset"
                    autoComplete="off"
                    {...field}
                  />
                )}
              />

              <Controller
                control={control}
                name="connectionParams.timeout"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      `Timeout for establishing connections, aka dial timeout. Value is in seconds. Optional, defaults to 0.`,
                    )}
                    type="number"
                    label="Timeout"
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
              <Controller
                control={control}
                name="connectionParams.readTimeout"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      `I/O read timeout. Value is in seconds. Optional, defaults to 0.`,
                    )}
                    type="number"
                    label="Read Timeout"
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
              <Controller
                control={control}
                name="connectionParams.writeTimeout"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      `I/O write timeout. Value is in seconds. Optional, defaults to 0.`,
                    )}
                    type="number"
                    label="Write Timeout"
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
            </Stack>
          </AccordionDetails>
        </Accordion>
      </CardContent>
    </Card>
  ) : null;
}
