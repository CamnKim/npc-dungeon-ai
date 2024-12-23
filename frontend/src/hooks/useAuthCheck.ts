import { useAuth0 } from '@auth0/auth0-react';
import { useContext } from 'react';
import { AuthContext } from '../contexts/auth.context';
import { useGetUserByAuth0 } from './queries';
import { User } from '../types/user.types';

export const useAuthCheck = (): { user: User | undefined; isLoading: boolean } | undefined => {
  const { isAuthenticated, user } = useAuth0();
  const { auth } = useContext(AuthContext);

  if (!isAuthenticated || !user?.sub) {
    return undefined;
  }

  if (auth.user) {
    return { user: auth.user, isLoading: false };
  }

  // Fetch user
  const { isLoading, data } = useGetUserByAuth0(user.sub);
  return { user: data, isLoading };
};
