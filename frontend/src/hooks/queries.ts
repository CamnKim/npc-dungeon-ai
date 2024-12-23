import { useAuth0 } from '@auth0/auth0-react';
import { useRef } from 'react';
import { NpcDungeonApi } from '../clients/npcDungeonApi';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { User, UserOnboard } from '../types/user.types';
import { ApiError } from '../types/error.types';

export const useNpcDungeonApi = () => {
  const { getAccessTokenSilently } = useAuth0();

  const npcDungeonApi = useRef(
    new NpcDungeonApi({
      origin: import.meta.env.VITE_NPC_DUNGEON_API_ORIGIN,
      getAccessToken: getAccessTokenSilently,
    })
  );

  return npcDungeonApi.current;
};

export const useGetUserByAuth0 = (sub: string, enabled: boolean = true) => {
  const npcDungeonApi = useNpcDungeonApi();
  return useQuery<User | undefined, ApiError>({
    queryKey: ['user', sub],
    queryFn: async () => npcDungeonApi.getUserByAuth0Sub(sub),
    enabled,
  });
};

export const useOnboardUser = () => {
  const queryClient = useQueryClient();
  const npcDungeonApi = useNpcDungeonApi();

  return useMutation<User, ApiError, UserOnboard>({
    mutationFn: async user => npcDungeonApi.onboardUser(user),
    onSuccess: data => {
      queryClient.setQueryData(['user', data.metadata.auth0Sub], data);
    },
  });
};
