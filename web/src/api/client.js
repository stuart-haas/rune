import axios from 'axios'
import { useToast } from '@/components/ui/toast'

const axiosClient = axios.create({
  baseURL: '/api',
})

axiosClient.interceptors.response.use(
  (response) => {
    const { toast } = useToast()
    const message = response.headers['x-message']
    if (message) {
      toast({
        title: 'Success',
        description: message,
        variant: 'default'
      })
    }
    return response
  },
  (error) => {
    const { toast } = useToast()
    const message = error.response?.headers?.['x-message']
    if (message) {
      toast({
        title: 'Error',
        description: message,
        variant: 'destructive'
      })
    }
    return Promise.reject(error)
  }
)

export default axiosClient