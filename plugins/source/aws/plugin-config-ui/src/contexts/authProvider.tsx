import React, { FC, useMemo } from 'react';

import { AuthContext } from './authContext';

interface Props {
  children: React.ReactNode;
  value: string;
}
export const AuthProvider: FC<Props> = ({ children, value }) => {
  const contextValue = useMemo(
    () => ({
      value,
    }),
    [value],
  );

  return <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>;
};
