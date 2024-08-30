import {
  useFormContext,
  ControlExclusiveToggle,
  ControlSecretField,
  Section,
  ControlTextField,
  ControlNumberField,
} from '@cloudquery/plugin-config-ui-lib';

export function FormConnectionFields() {
  const { watch } = useFormContext();
  const connectionType = watch('_connectionType');

  return (
    <Section
      title="Connect to your database"
      subtitle="Set up a connection to your MySQL instance."
    >
      <ControlExclusiveToggle
        name="_connectionType"
        options={[
          {
            label: 'Regular setup',
            value: 'fields',
          },
          {
            label: 'Connection string',
            value: 'string',
          },
        ]}
      />

      {connectionType === 'string' && (
        <ControlSecretField
          name="connectionString"
          helperText="Connection string to connect to the database. E.g. user:password@localhost:3306/dbname?tls=preferred\u0026readTimeout=1s\u0026writeTimeout=1s"
          label="Connection string"
        />
      )}
      {connectionType === 'fields' && (
        <>
          <ControlTextField
            name="host"
            helperText="Host to connect to. E.g. 1.2.3.4 or mydb.host.com. Optional, defaults to empty."
            label="Host"
          />
          <ControlNumberField
            name="port"
            helperText="Port to connect to. Optional, defaults to empty."
            label="Port"
          />
          <ControlTextField
            name="database"
            helperText="Name of the MySQL database you want to connect to. Optional, defaults to empty."
            label="Database"
          />
          <ControlTextField
            name="username"
            helperText="Username to use when authenticating. Optional, defaults to empty."
            label="Username"
          />
          <ControlSecretField
            name="password"
            label="Password"
            helperText="Password to use when authenticating. Optional, defaults to empty."
          />
        </>
      )}
    </Section>
  );
}
