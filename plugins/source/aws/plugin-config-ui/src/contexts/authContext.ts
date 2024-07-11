import { createContext } from 'react';

interface Context {
  value: string;
}

export const AuthContext = createContext<Context>({
  value: undefined as any,
});
