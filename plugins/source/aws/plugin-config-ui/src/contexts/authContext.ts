import { createContext } from 'react';

interface Context {
  token: string;
  team: string;
}

export const AuthContext = createContext<Context>({
  token: undefined as any,
  team: undefined as any,
});
