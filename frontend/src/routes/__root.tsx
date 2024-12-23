import { createRootRoute, Outlet } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import Navbar from '../components/navbar';
import { useMemo } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import UserProvider from '../components/UserProvider';

const Root = () => {
  const { isLoading: auth0Loading, user } = useAuth0();

  const sub = useMemo(() => {
    if (auth0Loading || !user) {
      return '';
    }
    return user.sub!;
  }, [auth0Loading, user]);

  if (!sub) {
    return (
      <>
        <Navbar />
        <hr />
        <Outlet />
        <TanStackRouterDevtools />
      </>
    );
  }

  return (
    <UserProvider auth0Sub={sub}>
      <Navbar />
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </UserProvider>
  );
};

export const Route = createRootRoute({
  component: Root,
});
