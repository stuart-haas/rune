import client from './client'
import { useQuery, useQueryClient, useMutation } from '@tanstack/vue-query'

const fetchDevices = async () => {
  const { data } = await client.get('/tailscale/devices')
  return data
}

const checkDevices = async () => {
  const { data } = await client.get('/tailscale/devices/check')
  return data
}

const syncDevices = async () => {
  const { data } = await client.post('/tailscale/devices/sync')
  return data
}

const syncDevice = async (id) => {
  const { data } = await client.post(`/tailscale/devices/${id}/sync`)
  return data
}

export const useTailscaleAPI = () => {
  const queryClient = useQueryClient()

  return {
    fetchDevices: () => useQuery({
      queryKey: ['tailscale-devices'],
      queryFn: fetchDevices,
    }),

    checkDevices: () => useQuery({
      queryKey: ['tailscale-devices-check'],
      queryFn: checkDevices,
    }),

    syncDevices: () => useMutation({
      mutationFn: syncDevices,
      onSuccess: () => {
        queryClient.invalidateQueries(['tailscale-devices', 'tailscale-devices-check'])
      }
    }),

    syncDevice: () => useMutation({
      mutationFn: syncDevice,
      onSuccess: () => {
        queryClient.invalidateQueries(['tailscale-devices', 'tailscale-devices-check'])
      }
    })
  }
}