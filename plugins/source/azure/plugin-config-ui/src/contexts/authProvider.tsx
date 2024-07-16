import React, { FC, useMemo } from 'react';

import { AuthContext } from './authContext';

interface Props {
  children: React.ReactNode;
  token: string;
  team: string;
}
export const AuthProvider: FC<Props> = ({ children, token, team }) => {
  const contextValue = useMemo(
    () => ({
      token,
      team,
    }),
    [token, team],
  );

  return <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>;
};
