import { Box, Button } from '@mui/material';

interface Props {
  handlePreviousStep?: () => void;
  handleCancel?: () => void;
  handleSubmit?: () => void;
  handleDelete?: () => void;
}

export function Footer({ handlePreviousStep, handleCancel, handleSubmit, handleDelete }: Props) {
  return (
    <Box display="flex" justifyContent="space-between">
      <Box display="flex" gap={2}>
        <Button onClick={handlePreviousStep} variant="text" color="secondary">
          Previous step
        </Button>
        {handleDelete && (
          <Button onClick={handleDelete} variant="contained" color="error">
            Delete
          </Button>
        )}
      </Box>
      <Box display="flex" gap={2}>
        <Button variant="text" onClick={handleCancel}>
          Cancel
        </Button>
        <Button type="submit" variant="contained">
          TODO:SUBMIT
        </Button>
      </Box>
    </Box>
  );
}
