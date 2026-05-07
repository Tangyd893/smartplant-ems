export interface User {
  id: number
  username: string
  real_name?: string
  role: string
  phone?: string
  email?: string
  status: number
}

export interface LoginData {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}