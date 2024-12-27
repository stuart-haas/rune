import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import client from './client'

const fetchOAuthClients = async () => {
  const { data } = await client.get('/oauth/clients')
  return data
}

const createOAuthClient = async (client) => {
  const { data } = await client.post('/oauth/clients', client)
  return data
}

const updateOAuthClient = async (client) => {
  const { data } = await client.put(`/oauth/clients/${client.id}`, client)
  return data
}

const deleteOAuthClient = async (id) => {
  const { data } = await client.delete(`/oauth/clients/${id}`)
  return data
}

export const useOAuthAPI = () => {
  const queryClient = useQueryClient()

  return {
    fetch: () => useQuery({
      queryKey: ['oauth-clients'],
      queryFn: fetchOAuthClients,
    }),

    create: () => useMutation({
      mutationFn: createOAuthClient,
      onSuccess: () => {
        queryClient.invalidateQueries(['oauth-clients'])
      }
    }),

    update: () => useMutation({
      mutationFn: updateOAuthClient,
      onSuccess: () => {
        queryClient.invalidateQueries(['oauth-clients'])
      }
    }),

    delete: () => useMutation({
      mutationFn: deleteOAuthClient,
      onSuccess: () => {
        queryClient.invalidateQueries(['oauth-clients'])
      }
    })
  }
}