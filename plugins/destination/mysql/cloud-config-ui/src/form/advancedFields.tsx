import { getFieldHelperText } from '@cloudquery/cloud-ui';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import Accordion from '@mui/material/Accordion';
import AccordionDetails from '@mui/material/AccordionDetails';
import AccordionSummary from '@mui/material/AccordionSummary';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import { Controller, useFormContext } from 'react-hook-form';

import { FormValues } from '../utils/formSchema';

export function AdvancedFields() {
  const { control } = useFormContext<FormValues>();

  return (
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
            <Typography variant="h6">Advanced Options</Typography>
          </AccordionSummary>
          <AccordionDetails sx={{ paddingLeft: 0 }}>
            <Stack spacing={2}>
              <Controller
                control={control}
                name="batchSize"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    required={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Maximum number of items that may be grouped together to be written in a single write. Default is 10,000.',
                    )}
                    label="Batch size"
                    {...field}
                  />
                )}
              />
              <Controller
                control={control}
                name="batchSizeBytes"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    required={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Maximum size of items that may be grouped together to be written in a single write. Default is 100,000,000 = 100MB.',
                    )}
                    label="Batch size (bytes)"
                    {...field}
                  />
                )}
              />
            </Stack>
          </AccordionDetails>
        </Accordion>
      </CardContent>
    </Card>
  );
}
