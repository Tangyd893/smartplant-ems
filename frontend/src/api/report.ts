import client, { ApiResponse } from './client'
import type { Report, ReportCreateData, ReportListParams } from './types/report'
import type { PageResult } from './types/device'

export const reportAPI = {
  list: (params?: ReportListParams) =>
    client.get<ApiResponse<PageResult<Report>>>('/report/list', {
      params,
    }).then(res => res.data),

  get: (id: number) =>
    client.get<ApiResponse<Report>>(`/report/${id}`).then(res => res.data),

  create: (data: ReportCreateData) =>
    client.post<ApiResponse<{ id: number }>>('/report', data).then(res => res.data),

  delete: (id: number) =>
    client.delete<ApiResponse<null>>(`/report/${id}`).then(res => res.data),

  download: (id: number) =>
    client.get(`/report/${id}/download`, {
      responseType: 'blob',
    }),
}