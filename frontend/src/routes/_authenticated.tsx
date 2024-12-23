import { useAuth0 } from '@auth0/auth0-react';
import { createFileRoute, Navigate, Outlet } from '@tanstack/react-router';

export const Route = createFileRoute('/_authenticated')({
  component: RouteComponent,
});

function RouteComponent() {
  const { isAuthenticated, isLoading } = useAuth0();
  if (isLoading) {
    return <h1>Loading...</h1>;
  }
  if (isAuthenticated) {
    return <Outlet />;
  }
  if (!isLoading && !isAuthenticated) {
    return <Navigate to="/login" />;
  }
  return <h1>shit out of luck</h1>;
}
