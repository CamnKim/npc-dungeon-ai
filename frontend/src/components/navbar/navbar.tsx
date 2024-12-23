import { useAuth0 } from '@auth0/auth0-react';
import { Link } from '@tanstack/react-router';

const Navbar = () => {
  const { isAuthenticated, logout } = useAuth0();
  return (
    <div className="p-2 flex gap-2">
      <Link to="/" className="[&.active]:font-bold">
        Home
      </Link>{' '}
      {isAuthenticated ? (
        <>
          <Link
            onClick={() => logout({ logoutParams: { returnTo: 'http://localhost:5173' } })}
            className="[&.active]:font-bold"
          >
            Logout
          </Link>{' '}
          <Link to="/profile" className="[&.active]:font-bold">
            Profile
          </Link>
        </>
      ) : (
        <Link to="/login" className="[&.active]:font-bold">
          Login
        </Link>
      )}
    </div>
  );
};

export default Navbar;
