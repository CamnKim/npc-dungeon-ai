export type UserOnboard = {
  username: string;
  email: string;
  auth0Sub: string;
  region?: string;
  dateOfBirth?: string;
};

export type User = {
  id: string;
  username: string;
  email: string;
  createdAt: string;
  updatedAt: string;
  region?: string;
  dateOfBirth?: string;
  isActive: boolean;
  metadata: {
    auth0Sub: string;
  };
};

export type UserRaw = {
  id: string;
  auth0_sub: string;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
  region?: string | null;
  date_of_birth?: string | null;
  is_active: boolean;
};

export const UserRawToUser = (userRaw: UserRaw): User => ({
  id: userRaw.id,
  username: userRaw.username,
  email: userRaw.email,
  createdAt: userRaw.created_at,
  updatedAt: userRaw.updated_at,
  region: userRaw.region || undefined,
  dateOfBirth: userRaw.date_of_birth || undefined,
  isActive: userRaw.is_active,
  metadata: {
    auth0Sub: userRaw.auth0_sub,
  },
});
