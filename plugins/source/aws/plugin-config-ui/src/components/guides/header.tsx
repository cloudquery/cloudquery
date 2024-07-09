import { Link as LinkIcon } from '@mui/icons-material';
import { Box, Button, Typography } from '@mui/material';

export function AWSGuideHeader() {
  return (
    <Box display="flex" justifyContent="space-between" alignItems="center">
      <Typography variant="h6">Setup guide</Typography>
      <Button
        color="secondary"
        variant="outlined"
        endIcon={<LinkIcon />}
        onClick={() => {
          //TODO
        }}
      >
        Open docs
      </Button>
    </Box>
  );
}
