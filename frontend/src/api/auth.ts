import client, { ApiResponse } from './client'
import type { LoginData, LoginResponse } from './types/auth'

export const authAPI = {
  login: (data: LoginData) =>
    client.post<ApiResponse<LoginResponse>>('/auth/login', data).then(res => res.data),

  logout: () =>
    client.post<ApiResponse<null>>('/auth/logout').then(res => res.data),

  getCurrentUser: () =>
    client.get<ApiResponse<LoginResponse['user']>>('/auth/me').then(res => res.data),
}