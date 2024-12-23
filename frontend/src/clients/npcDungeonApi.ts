import { ApiError } from '../types/error.types';
import { User, UserOnboard, UserRaw, UserRawToUser } from '../types/user.types';

type ApiConfig = {
  origin: string;
  getAccessToken: () => Promise<string>;
};

type OnboardReqBody = {
  username: string;
  email: string;
  region?: string;
  auth0_sub: string;
  date_of_birth?: string;
};

export class NpcDungeonApi {
  private readonly config: ApiConfig;

  constructor(config: ApiConfig) {
    this.config = config;
  }

  headers(token: string) {
    return {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    };
  }

  async getUserByAuth0Sub(auth0Sub: string): Promise<User | undefined> {
    try {
      const token = await this.config.getAccessToken();
      const response = await fetch(`${this.config.origin}/auth/users/${auth0Sub}`, {
        headers: this.headers(token),
        method: 'GET',
      });

      if (response.status === 404) {
        return undefined;
      }
      if (!response.ok) {
        throw new ApiError('Failed to get user by Auth0 sub', {
          status: response.status,
        });
      }
      // Convert the raw user to a User object
      const res = (await response.json()) as UserRaw;
      return UserRawToUser(res);
    } catch (error) {
      if (error instanceof ApiError) {
        throw error;
      }
      throw new ApiError('Failed to get user by Auth0 sub', {
        error,
      });
    }
  }

  async onboardUser(user: UserOnboard): Promise<User> {
    try {
      const token = await this.config.getAccessToken();

      const body: OnboardReqBody = {
        ...user,
        auth0_sub: user.auth0Sub,
        date_of_birth: user.dateOfBirth,
      };
      const response = await fetch(`${this.config.origin}/users`, {
        headers: this.headers(token),
        method: 'POST',
        body: JSON.stringify(body),
      });

      // Convert the raw user to a User object
      const res = (await response.json()) as UserRaw;
      return UserRawToUser(res);
    } catch (error) {
      throw new ApiError('Failed to onboard user', {
        error,
      });
    }
  }
}
