import { PropsWithChildren, useContext, useEffect, useMemo } from 'react';
import { AuthContext } from '../../contexts/auth.context';
import { useGetUserByAuth0 } from '../../hooks/queries';
import { useAuth0 } from '@auth0/auth0-react';
import { useNavigate } from '@tanstack/react-router';

type Props = {
  auth0Sub: string;
};

const UserProvider = ({ auth0Sub, children }: PropsWithChildren<Props>) => {
  const { isAuthenticated, isLoading: authIsLoading } = useAuth0();
  const { auth, setAuthValue } = useContext(AuthContext);

  const { data, isLoading } = useGetUserByAuth0(auth0Sub);

  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && data && !auth.user) {
      setAuthValue({ ...auth, user: data });
    }
  }, [isLoading, data, setAuthValue, auth]);

  const shouldOnboard = useMemo(() => {
    return !data && isAuthenticated && !isLoading && !authIsLoading;
  }, [data, isAuthenticated, isLoading, authIsLoading]);

  if (shouldOnboard) {
    navigate({ to: '/onboard' });
  }

  if (!data || isLoading) {
    return <h1>Loading...</h1>;
  }
  return children;
};

export default UserProvider;
