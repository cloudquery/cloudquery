import {
  useFormContext,
  ControlDateTimeField,
  ControlBooleanField,
} from '@cloudquery/plugin-config-ui-lib';

export function StartTime() {
  const { watch } = useFormContext();
  const startTimeEnabled = watch('_startTimeEnabled');

  return (
    <ControlDateTimeField
      name="start_time"
      label="Start time"
      helperText="The earliest news date that the source should fetch."
      disabled={!startTimeEnabled}
      InputProps={{
        endAdornment: <ControlBooleanField name="_startTimeEnabled" type="toggle" label="" />,
      }}
    />
  );
}
