interface Props {}

export function SelectServices({}: Props) {
  return (
    <>select services</>
    // <FormFieldGroup title="PostgreSQL Connection">
    //   <Controller
    //     control={control}
    //     name="spec.host"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Host"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.port"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Port"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.database"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Database"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.username"
    //     render={({ field, fieldState }) => (
    //       <Stack direction="row" spacing={2}>
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="Username"
    //           {...field}
    //           disabled={defaultUsername === '${username}' && !usernameResetted}
    //           value={
    //             defaultUsername === '${username}' && !usernameResetted
    //               ? envPlaceholder
    //               : field.value
    //           }
    //         />
    //         {defaultUsername === '${username}' && (
    //           <FormFieldReset
    //             isResetted={usernameResetted}
    //             inputSelectorToFocus='input[name="spec.username"]'
    //             onCancel={() => handelCancelReset('username')}
    //             onReset={() => handleReset('username')}
    //           />
    //         )}
    //       </Stack>
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.password"
    //     render={({ field, fieldState }) => (
    //       <Stack direction="row" spacing={2}>
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="Password"
    //           {...field}
    //           disabled={defaultPassword === '${password}' && !passwordResetted}
    //           value={
    //             defaultPassword === '${password}' && !passwordResetted
    //               ? envPlaceholder
    //               : field.value
    //           }
    //         />
    //         {defaultPassword === '${password}' && (
    //           <FormFieldReset
    //             isResetted={passwordResetted}
    //             inputSelectorToFocus='input[name="spec.password"]'
    //             onCancel={() => handelCancelReset('password')}
    //             onReset={() => handleReset('password')}
    //           />
    //         )}
    //       </Stack>
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.clientEncoding"
    //     render={({ field, fieldState }) => (
    //       <TextField
    //         error={!!fieldState.error}
    //         fullWidth={true}
    //         helperText={fieldState.error?.message}
    //         label="Client Encoding"
    //         {...field}
    //       />
    //     )}
    //   />
    //   <Controller
    //     control={control}
    //     name="spec.ssl"
    //     render={({ field }) => <FormControlLabel control={<Switch {...field} />} label="SSL" />}
    //   />
    //   {sslValue && (
    //     <Controller
    //       control={control}
    //       name="spec.sslMode"
    //       render={({ field, fieldState }) => (
    //         <TextField
    //           error={!!fieldState.error}
    //           fullWidth={true}
    //           helperText={fieldState.error?.message}
    //           label="SSL Mode"
    //           select={true}
    //           {...field}
    //         >
    //           <MenuItem value={''} hidden={true} />
    //           {sslModeValues.map((value) => (
    //             <MenuItem key={value} value={value}>
    //               {value}
    //             </MenuItem>
    //           ))}
    //         </TextField>
    //       )}
    //     />
    //   )}
    // </FormFieldGroup>
  );
}
