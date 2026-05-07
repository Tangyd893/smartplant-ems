import client, { ApiResponse } from './client'
import type { EnergyRecord, EnergyStats, EnergyListParams } from './types/energy'
import type { PageResult } from './types/device'

export const energyAPI = {
  list: (params?: EnergyListParams) =>
    client.get<ApiResponse<PageResult<EnergyRecord>>>(
      '/energy/list',
      { params }
    ).then(res => res.data),

  get: (id: number) =>
    client.get<ApiResponse<EnergyRecord>>(`/energy/${id}`).then(res => res.data),

  create: (data: Partial<EnergyRecord>) =>
    client.post<ApiResponse<{ id: number }>>('/energy', data).then(res => res.data),

  update: (id: number, data: Partial<EnergyRecord>) =>
    client.put<ApiResponse<null>>(`/energy/${id}`, data).then(res => res.data),

  getStats: (params: { device_id?: number; start_time?: string; end_time?: string }) =>
    client.get<ApiResponse<EnergyStats[]>>('/energy/stats', { params }).then(res => res.data),

  getDeviceStats: (deviceId: number, startTime?: string, endTime?: string) =>
    client.get<ApiResponse<EnergyStats>>(
      `/energy/device/${deviceId}/stats`,
      {
        params: { start_time: startTime, end_time: endTime },
      }
    ).then(res => res.data),
}