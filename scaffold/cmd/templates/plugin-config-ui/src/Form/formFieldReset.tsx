import { useCallback } from 'react';

import Button from '@mui/material/Button';

import { UseFormSetValue } from 'react-hook-form';
import { FormValues, existingSecretValue } from '../utils/formSchema';

interface Props {
  initialValue: string | symbol | undefined;
  isReadonly: boolean;
  path: string;
  setValue: UseFormSetValue<FormValues>;
}

export function FormFieldReset({ initialValue, isReadonly, path, setValue }: Props) {
  const handleReset = useCallback(() => {
    setValue(path as any, '', { shouldDirty: true });
    setTimeout(() => {
      const element = document.querySelector(`[name="${path}"]`) as HTMLInputElement | null;
      element?.focus();
    }, 0);
  }, [path, setValue]);

  const handleCancel = useCallback(() => {
    setValue(path as any, existingSecretValue as any, { shouldDirty: true });
  }, [path, setValue]);

  return (
    <>
      {isReadonly && (
        <Button onClick={handleReset} sx={{ width: 100 }} variant="outlined">
          Reset
        </Button>
      )}
      {!isReadonly && initialValue === existingSecretValue && (
        <Button onClick={handleCancel} sx={{ width: 100 }} variant="outlined">
          Cancel
        </Button>
      )}
    </>
  );
}
