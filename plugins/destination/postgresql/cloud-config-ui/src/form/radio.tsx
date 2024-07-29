import { ReactNode } from 'react';

import Radio from '@mui/material/Radio';
import Stack from '@mui/material/Stack';
import useTheme from '@mui/material/styles/useTheme';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import Typography from '@mui/material/Typography';

interface Props<Value extends string | number | boolean> {
  /** Callback that is called when the selected value changes. */
  onChange: (newValue: Value) => void;
  /** The currently selected value. */
  value: Value;
  /** The title of the radio group. */
  title?: string;
  /** The radio buttons to display. */
  items: Array<{
    title: ReactNode;
    subtitle?: ReactNode;
    value: string | number | boolean;
    disabled?: boolean;
  }>;
}

/**
 * This component displays a group of radio buttons.
 *
 * @public
 */
export function RadioGroupSelector<Value extends string | number | boolean>({
  onChange,
  title,
  value,
  items,
}: Props<Value>) {
  const { palette } = useTheme();

  return (
    <ToggleButtonGroup
      aria-label={title}
      color="primary"
      exclusive={true}
      onChange={(event, newValue) => {
        if (event.type === 'click' && newValue !== null) {
          onChange(newValue);
        }
      }}
      value={value}
    >
      <Stack direction="row" spacing={2} width="100%">
        {items.map((item) => {
          const isSelected = value === item.value;

          return (
            <ToggleButton
              key={String(item.value)}
              disabled={item.disabled}
              fullWidth={true}
              value={item.value}
              sx={{
                padding: item.subtitle ? undefined : '2px',
              }}
            >
              <Radio checked={isSelected} />
              <Stack marginLeft={0.5} paddingY={1.25} spacing={0.5}>
                <Typography
                  color={isSelected ? palette.text.primary : palette.text.secondary}
                  sx={{ opacity: isSelected ? 1 : 0.8 }}
                  variant="body1Bold"
                >
                  {item.title}
                </Typography>
                {!!item.subtitle && (
                  <Typography
                    color={isSelected ? palette.text.primary : palette.text.secondary}
                    sx={{ opacity: isSelected ? 1 : 0.8 }}
                    variant="body2"
                  >
                    {item.subtitle}
                  </Typography>
                )}
              </Stack>
            </ToggleButton>
          );
        })}
      </Stack>
    </ToggleButtonGroup>
  );
}
