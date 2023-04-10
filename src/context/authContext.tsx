import React, { createContext, ReactNode } from 'react';
import { useState } from 'react';
import { GetCookie } from '../function/cookies';

type Props = {
  children?: ReactNode;
};

type iAuthContext = {
  authenticated: boolean;
  setAuthenticated: (newState: boolean) => void;
};

const initialValue = {
  authenticated: GetCookie('token') ? true : false,
  setAuthenticated: () => {},
};

export const AuthContext = createContext<iAuthContext>(initialValue);

export const AuthProvider = ({ children }: Props): JSX.Element => {
  const [authenticated, setAuthenticated] = useState(
    initialValue.authenticated,
  );
  return (
    <AuthContext.Provider value={{ authenticated, setAuthenticated }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
