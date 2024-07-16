import Radio from '@mui/material/Radio';
import Stack from '@mui/material/Stack';
import useTheme from '@mui/material/styles/useTheme';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import Typography from '@mui/material/Typography';
import { forwardRef } from 'react';

type Option = {
  value: any;
  label: string;
};

interface Props {
  options: Option[];
  onChange: (newValue: boolean) => void;
  value: boolean;
}

export const ExclusiveToggle = forwardRef(function ExclusiveToggle(
  { options, onChange, value }: Props,
  ref,
) {
  const { palette } = useTheme();

  const buttonSx = {
    padding: 0.5,
  };

  return (
    <Stack gap={4}>
      <ToggleButtonGroup
        aria-label="Exclusive toggle"
        color="primary"
        exclusive={true}
        onChange={(_, newValue) => {
          if (!!newValue) {
            onChange(newValue);
          }
        }}
        value={value}
        ref={ref}
      >
        <Stack direction="row" spacing={2} width="100%">
          {options.map((option) => (
            <ToggleButton key={option.value} sx={buttonSx} fullWidth={true} value={option.value}>
              <Radio checked={value === option.value}></Radio>
              <Stack marginLeft={0.5} paddingY={1.25} spacing={0.5}>
                <Typography
                  color={value === option.value ? palette.text.primary : palette.text.secondary}
                  sx={{ opacity: value === option.value ? 1 : 0.8 }}
                  variant="body1"
                >
                  {option.label}
                </Typography>
              </Stack>
            </ToggleButton>
          ))}
        </Stack>
      </ToggleButtonGroup>
    </Stack>
  );
});
