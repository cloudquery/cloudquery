import React from 'react';

import FormControl from '@mui/material/FormControl';
import FormHelperText from '@mui/material/FormHelperText';

import { useGCPConnector } from '../context/GCPConnectorContext';

export function AuthError() {
  const { authenticationError } = useGCPConnector();

  return (
    authenticationError && (
      <FormControl>
        {<FormHelperText error={true}>Network error: {authenticationError.message}</FormHelperText>}
      </FormControl>
    )
  );
}
