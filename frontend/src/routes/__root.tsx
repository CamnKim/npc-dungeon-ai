import { createRootRoute, Outlet, useNavigate } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import Navbar from '../components/navbar';
import { useContext, useEffect, useMemo } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { AuthContext } from '../contexts/auth.context';
import { useGetUserByAuth0 } from '../hooks/queries';

const Root = () => {
  const { isLoading: auth0Loading, isAuthenticated, user } = useAuth0();
  const { setAuthValue } = useContext(AuthContext);

  const sub = useMemo(() => {
    if (auth0Loading || !user) {
      return '';
    }
    return user.sub!;
  }, [auth0Loading, user]);

  const { data, isLoading } = useGetUserByAuth0(sub, auth0Loading);

  const navigate = useNavigate();

  useEffect(() => {
    console.log(data);
    if (!auth0Loading && !isLoading && isAuthenticated) {
      if (!data) {
        console.log('User not found, onboard user');
        navigate({ to: '/onboard' });
      }
      setAuthValue({ user: data, auth0: user });
    }
  }, [auth0Loading, isAuthenticated, data, isLoading, setAuthValue, user]);

  return (
    <>
      <Navbar />
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  );
};

export const Route = createRootRoute({
  component: Root,
});
