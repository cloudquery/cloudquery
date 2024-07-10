import { Box, useTheme } from '@mui/material';

type Props = {
  width?: number;
  height?: number;
  src: string;
  alt?: string;
};

const PADDING = 4;

export function Logo({ width = 24, height = 24, src, alt }: Props) {
  const { palette } = useTheme();
  return (
    <Box
      sx={{
        borderRadius: `${PADDING}px`,
        backgroundColor: palette.secondary.light,
        height,
        width,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <img src={src} alt={alt ?? src} height={height - PADDING} width={width - PADDING} />
    </Box>
  );
}
