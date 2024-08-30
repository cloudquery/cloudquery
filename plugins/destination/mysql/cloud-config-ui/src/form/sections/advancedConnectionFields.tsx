import {
  useFormContext,
  CollapsibleSection,
  ControlTextField,
  ControlBooleanField,
  ControlSelectField,
  ControlNumberField,
} from '@cloudquery/plugin-config-ui-lib';

import { tlsModeValues } from '../../utils/formSchema';

export function AdvancedConnectionFields() {
  const { watch } = useFormContext();

  const tlsEnabled = watch('connectionParams.tls');
  const parseTimeEnabled = watch('connectionParams.parseTime');
  const connectionType = watch('_connectionType');

  return connectionType === 'fields' ? (
    <CollapsibleSection defaultExpanded={false} title="Advanced Connection Options">
      <ControlBooleanField
        name="tcp"
        label="TCP"
        type="toggle"
        helperText="If true, will enable connection over TCP to the server. Optional, defaults to true."
      />
      <ControlBooleanField
        name="connectionParams.tls"
        label="TLS"
        type="toggle"
        helperText="If true, will enabled TLS/SSL encrypted connection to the server. Optional, defaults to false."
      />

      {tlsEnabled && (
        <ControlSelectField
          name="connectionParams.tlsMode"
          helperText="SSL connections to encrypt client/server communications using TLS protocols for increased security."
          label="TLS Mode"
          options={[...tlsModeValues]}
        />
      )}
      <ControlBooleanField
        name="connectionParams.parseTime"
        label="Parse Time"
        type="toggle"
        helperText="If true, changes the output type of DATE and DATETIME values to time.Time instead of []byte / string. Optional, defaults to false."
      />
      {parseTimeEnabled && (
        <ControlTextField
          name="connectionParams.loc"
          helperText={`Sets the location for time.Time values. "Local" sets the system's location. Optional, defaults to UTC.`}
          label="Location"
        />
      )}
      <ControlTextField
        name="connectionParams.charset"
        helperText="Sets the charset used for client-server interaction. Multiple charsets can be configured with comma separation (ex. utf8mb4,utf8). Optional, defaults to utf8mb4."
        label="Charset"
      />
      <ControlNumberField
        name="connectionParams.timeout"
        helperText={`Timeout for establishing connections, aka dial timeout. Value is in seconds. Optional, defaults to 0.`}
        label="Timeout"
      />
      <ControlNumberField
        name="connectionParams.readTimeout"
        helperText={`I/O read timeout. Value is in seconds. Optional, defaults to 0.`}
        label="Read Timeout"
      />
      <ControlNumberField
        name="connectionParams.writeTimeout"
        helperText={`I/O write timeout. Value is in seconds. Optional, defaults to 0.`}
        label="Write Timeout"
      />
    </CollapsibleSection>
  ) : null;
}
