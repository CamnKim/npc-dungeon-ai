import { useContext } from 'react';
import { AuthContext } from '../contexts/auth.context';

const ProfilePage = () => {
  const {
    auth: { user },
  } = useContext(AuthContext);
  return (
    <div>
      <h1>Profile Page</h1>
      <h2>{user?.username}</h2>
    </div>
  );
};

export default ProfilePage;
