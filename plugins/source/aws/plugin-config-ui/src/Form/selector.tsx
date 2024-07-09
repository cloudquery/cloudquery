import React, { useEffect } from 'react';

import Radio from '@mui/material/Radio';
import Stack from '@mui/material/Stack';
import useTheme from '@mui/material/styles/useTheme';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import Typography from '@mui/material/Typography';

type Option = {
  value: any;
  label: string;
};

interface Props {
  optionA: Option;
  optionB: Option;
  onChange: (newValue: boolean) => void;
  value: boolean;
}

export function ExclusiveToggle({ optionA, optionB, onChange, value }: Props) {
  const { palette } = useTheme();

  const buttonSx = {
    padding: 0.5,
  };

  return (
    <Stack gap={4}>
      <ToggleButtonGroup
        aria-label="Use Existing"
        color="primary"
        exclusive={true}
        onChange={(_, newValue) => {
          onChange(newValue);
        }}
        value={value}
      >
        <Stack direction="row" spacing={2} width="100%">
          <ToggleButton sx={buttonSx} fullWidth={true} value={optionA.value}>
            <Radio checked={value === optionA.value}></Radio>
            <Stack marginLeft={0.5} paddingY={1.25} spacing={0.5}>
              <Typography
                color={value === optionA.value ? palette.text.primary : palette.text.secondary}
                sx={{ opacity: value === optionA.value ? 1 : 0.8 }}
                variant="body1"
              >
                {optionA.label}
              </Typography>
            </Stack>
          </ToggleButton>
          <ToggleButton sx={buttonSx} fullWidth={true} value={optionB.value}>
            <Radio checked={value === optionB.value}></Radio>
            <Stack marginLeft={0.5} paddingY={1.25} spacing={0.5}>
              <Typography
                color={value === optionB.value ? palette.text.primary : palette.text.secondary}
                sx={{ opacity: value === optionB.value ? 1 : 0.8 }}
                variant="body1"
              >
                {optionB.label}
              </Typography>
            </Stack>
          </ToggleButton>
        </Stack>
      </ToggleButtonGroup>
    </Stack>
  );
}
