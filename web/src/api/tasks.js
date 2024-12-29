import client from './client'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'

const fetchTasks = async () => {
  const { data } = await client.get('/tasks')
  return data
}

const createTask = async (task) => {
  const { data } = await client.post('/tasks', task)
  return data
}

const updateTask = async (task) => {
  const { data } = await client.put(`/tasks/${task.ID}`, task)
  return data
}

const deleteTask = async (id) => {
  const { data } = await client.delete(`/tasks/${id}`)
  return data
}

export const useTasksAPI = () => {
  const queryClient = useQueryClient()
  return {
    fetch: () => useQuery({ queryKey: ['tasks'], queryFn: fetchTasks }),
    create: () => useMutation({ mutationFn: createTask, onSuccess: () => queryClient.invalidateQueries(['tasks']) }),
    update: () => useMutation({ mutationFn: updateTask, onSuccess: () => queryClient.invalidateQueries(['tasks']) }),
    delete: () => useMutation({ mutationFn: deleteTask, onSuccess: () => queryClient.invalidateQueries(['tasks']) }),
  }
}