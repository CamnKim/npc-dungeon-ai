import { StrictMode } from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import { Auth0Provider, useAuth0 } from '@auth0/auth0-react';

// Import the generated route tree
import { routeTree } from './routeTree.gen';
import { AuthProvider } from './contexts/auth.context';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

// Create a new router instance
const router = createRouter({
  routeTree,
  context: {
    auth: undefined!,
  },
});

// Create query client
const queryClient = new QueryClient();

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

const InnerApp = () => {
  const auth = useAuth0();
  return <RouterProvider router={router} context={{ auth }} />;
};

// Render the app
const rootElement = document.getElementById('root')!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <Auth0Provider
      domain={import.meta.env.VITE_AUTH0_DOMAIN}
      clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
      useRefreshTokens
      cacheLocation="localstorage"
      authorizationParams={{
        redirect_uri: 'http://localhost:5173',
        audience: 'https://dev-t65wdsllc40f4atm.us.auth0.com/api/v2/',
        scope: 'read:current_user update:current_user_metadata openid profile email',
      }}
    >
      <StrictMode>
        <QueryClientProvider client={queryClient}>
          <AuthProvider>
            <InnerApp />
          </AuthProvider>
        </QueryClientProvider>
      </StrictMode>
    </Auth0Provider>
  );
}
