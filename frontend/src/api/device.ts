import client, { ApiResponse } from './client'
import type {
  Device,
  DeviceListParams,
  DeviceCreateData,
  DeviceUpdateData,
  PageResult,
} from './types/device'

export const deviceAPI = {
  list: (params?: DeviceListParams) =>
    client.get<ApiResponse<PageResult<Device>>>('/device/list', { params }).then(res => res.data),

  get: (id: number) =>
    client.get<ApiResponse<Device>>(`/device/${id}`).then(res => res.data),

  create: (data: DeviceCreateData) =>
    client.post<ApiResponse<{ id: number }>>('/device', data).then(res => res.data),

  update: (id: number, data: DeviceUpdateData) =>
    client.put<ApiResponse<null>>(`/device/${id}`, data).then(res => res.data),

  delete: (id: number) =>
    client.delete<ApiResponse<null>>(`/device/${id}`).then(res => res.data),

  getRealTimeData: (id: number) =>
    client.get<ApiResponse<Record<string, number | string>>>(
      `/device/${id}/realtime`
    ).then(res => res.data),
}