import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import client from './client'


const fetchNodes = async () => {
  const { data } = await client.get('/nodes')
  return data
}

const fetchNodeTags = async () => {
  const { data } = await client.get('/nodes/tags')
  return data
}

const createNode = async (node) => {
  const { data } = await client.post('/nodes', node)
  return data
}

const updateNode = async (node) => {
  const { data } = await client.put(`/nodes/${node.ID}`, node)
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

    fetchTags: () => useQuery({
      queryKey: ['node-tags'],
      queryFn: fetchNodeTags,
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

    delete: () => useMutation({
      mutationFn: deleteNode,
      onSuccess: () => {
        queryClient.invalidateQueries(['nodes'])
      }
    })
  }
}