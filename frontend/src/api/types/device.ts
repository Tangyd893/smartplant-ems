export interface Device {
  id: number
  device_code: string
  device_name: string
  device_type: string
  location: string
  status: number
  power_rating: number
  created_at?: string
  updated_at?: string
}

export interface DeviceListParams {
  page?: number
  size?: number
  device_type?: string
  status?: number
}

export interface DeviceCreateData {
  device_code: string
  device_name: string
  device_type: string
  location?: string
  power_rating?: number
}

export interface DeviceUpdateData {
  id: number
  device_name?: string
  location?: string
  status?: number
  power_rating?: number
}

export interface PageResult<T> {
  records: T[]
  total: number
  page: number
  size: number
}