import { User as Auth0User, useAuth0 } from '@auth0/auth0-react';
import { createContext, PropsWithChildren, useMemo, useState } from 'react';
import { User } from '../types/user.types';

type Auth = {
  auth0?: Auth0User;
  user?: User;
};

export type AuthContextType = {
  auth: Auth;
  setAuthValue: (value: Auth) => void;
};

export const AuthContext = createContext<AuthContextType>({ auth: {}, setAuthValue: () => {} });

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const { user } = useAuth0();
  const [authValue, setAuthValue] = useState<Auth>({ auth0: user });

  const value = useMemo<AuthContextType>(() => ({ auth: authValue, setAuthValue }), [authValue]);
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
