import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import client from './client'


const fetchNodes = async () => {
  const { data } = await client.get('/nodes')
  return data
}

const createNode = async (node) => {
  const { data } = await client.post('/nodes', node)
  return data
}

const updateNode = async (node) => {
  const { data } = await client.put(`/nodes/${node.id}`, node)
  return data
}

const deleteNode = async (id) => {
  const { data } = await client.delete(`/nodes/${id}`)
  return data
}

export const useNodesAPI = () => {
  const queryClient = useQueryClient()
  
  return {
    fetch: () => useQuery({
      queryKey: ['nodes'],
      queryFn: fetchNodes,
    }),

    create: () => useMutation({
      mutationFn: createNode,
      onSuccess: () => {
        queryClient.invalidateQueries(['nodes'])
      }
    }),

    update: () => useMutation({
      mutationFn: updateNode,
      onSuccess: () => {
        queryClient.invalidateQueries(['nodes'])
      }
    }),

    remove: () => useMutation({
      mutationFn: deleteNode,
      onSuccess: () => {
        queryClient.invalidateQueries(['nodes'])
      }
    })
  }
}