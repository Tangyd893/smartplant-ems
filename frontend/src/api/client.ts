import axios, { AxiosInstance, AxiosError, InternalAxiosRequestConfig, AxiosResponse } from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api'

export interface ApiResponse<T> {
  code: number
  msg: string
  data: T
}

const client: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

client.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('token')
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error: AxiosError) => Promise.reject(error)
)

client.interceptors.response.use(
  (response: AxiosResponse) => response,
  (error: AxiosError<{ msg?: string }>) => {
    const { response } = error

    if (response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
      return Promise.reject(new Error('登录已过期，请重新登录'))
    }

    const message = response?.data?.msg || error.message || '网络请求失败'
    return Promise.reject(new Error(message))
  }
)

export const request = async <T = unknown>(config: Parameters<typeof client.request>[0]): Promise<AxiosResponse<T>> => {
  return client.request(config) as unknown as Promise<AxiosResponse<T>>
}

export default client