import { useState } from 'react';
import { useOnboardUser } from '../hooks/queries';
import { useAuth0 } from '@auth0/auth0-react';

type FormData = {
  username: string;
  region?: string;
  dob?: string;
};

const OnboardingPage = () => {
  const [formData, setFormData] = useState<FormData>({ username: '' });
  console.log(formData);

  const { user } = useAuth0();

  const { mutateAsync } = useOnboardUser();

  if (!user || !user.email || !user.sub || user.email === undefined || user.sub === undefined) {
    return <h1>Loading...</h1>;
  }

  const submit = () => {
    mutateAsync({
      username: formData.username,
      email: user.email!,
      auth0Sub: user.sub!,
      region: formData.region,
      dateOfBirth: formData.dob,
    });
  };

  return (
    <div>
      <h1>Onboarding!</h1>
      <label htmlFor="username">Username: </label>
      <input
        type="text"
        placeholder="Username"
        id="username"
        value={formData?.username}
        onChange={event => setFormData({ ...formData, username: event.target.value })}
      />
      <br />
      <label htmlFor="region">Region (optional): </label>
      <input
        type="text"
        placeholder="Region"
        id="region"
        value={formData?.region}
        onChange={event => setFormData({ ...formData, region: event.target.value })}
      />
      <br />
      <label htmlFor="dob">Date of Birth (optional): </label>
      <input
        type="date"
        id="dob"
        value={formData?.dob}
        onChange={event => setFormData({ ...formData, dob: event.target.value })}
      />
      <button onClick={submit}>Submit</button>
    </div>
  );
};

export default OnboardingPage;
