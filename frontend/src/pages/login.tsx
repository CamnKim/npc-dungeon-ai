import { useAuth0 } from '@auth0/auth0-react';

const LoginPage = () => {
  const { loginWithRedirect, logout, isAuthenticated } = useAuth0();

  return (
    <div>
      <h1>Login</h1>
      {isAuthenticated ? (
        <button onClick={() => logout({ logoutParams: { returnTo: 'http://localhost:5173' } })}>Log Out</button>
      ) : (
        <button onClick={() => loginWithRedirect()}>Login</button>
      )}
    </div>
  );
};

export default LoginPage;
