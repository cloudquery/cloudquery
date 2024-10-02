import React, { ChangeEvent, useRef, useState } from 'react';

import { Controller, getFieldHelperText, useFormContext } from '@cloudquery/plugin-config-ui-lib';
import CheckIcon from '@mui/icons-material/Check';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import FormControl from '@mui/material/FormControl';
import FormHelperText from '@mui/material/FormHelperText';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

// TODO: candidate for generic and in the lib
export function UploadJSON({ name = 'service_account_key_json' }: { name: string }) {
  const [error, setError] = useState<boolean>(false);
  const { watch, setValue } = useFormContext();
  const jsonFile = watch(name);

  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileUpload = async (event: ChangeEvent<HTMLInputElement>) => {
    if (!event.target.files) {
      return;
    }

    const file = event.target.files[0];
    if (file) {
      try {
        const fileText = await file.text();
        // test parse, error if fail
        JSON.parse(fileText);

        setValue(name, fileText);
        setError(false);
      } catch {
        setError(true);
      }
    }
  };

  return (
    <Stack gap={1.5}>
      <Box>
        <Controller
          name={name}
          render={({ fieldState }) => (
            <FormControl>
              <Button
                size="large"
                variant="contained"
                component="label"
                fullWidth={false}
                endIcon={jsonFile && <CheckIcon />}
              >
                {jsonFile ? 'JSON file uploaded successfully' : 'Upload JSON file'}
                <input
                  ref={fileInputRef}
                  accept="application/JSON"
                  hidden={true}
                  onChange={handleFileUpload}
                  type="file"
                />
              </Button>
              {(!!fieldState.error?.message || error) && (
                <FormHelperText error={true}>
                  {error
                    ? `JSON file could not be parsed.`
                    : getFieldHelperText(fieldState.error?.message, '')}
                </FormHelperText>
              )}
            </FormControl>
          )}
        />
      </Box>
      {jsonFile ? (
        <Typography variant="body1">Reupload your Service Account JSON file</Typography>
      ) : (
        <Typography variant="body1" fontWeight="bold">
          Upload your Service Account JSON file
        </Typography>
      )}
    </Stack>
  );
}
