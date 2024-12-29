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
    if (error.response.data.error) {
      toast({
        title: 'Error',
        description: error.response.data.error,
        variant: 'destructive'
      })
    }
    return Promise.reject(error)
  }
)

export default axiosClient