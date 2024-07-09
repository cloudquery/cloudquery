import Box from '@mui/material/Box';

interface Props {
  text: string;
}

export function CodeSnippet({ text }: Props) {
  return (
    <Box
      borderRadius="10px"
      color="#F97583"
      component="code"
      display="block"
      fontSize="14px"
      lineHeight={1.5}
      padding={2}
      position="relative"
      sx={{ WebkitBoxDecorationBreak: 'clone' }}
      textAlign="left"
      whiteSpace="pre-line"
    >
      {text}
    </Box>
  );
}
