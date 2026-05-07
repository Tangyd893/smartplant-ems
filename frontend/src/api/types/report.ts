export interface Report {
  id: number
  report_name: string
  report_type: string
  period_type: string
  period_start: string
  period_end: string
  created_by: number
  creator_name?: string
  file_path?: string
  status: number
  created_at?: string
  completed_at?: string
}

export interface ReportCreateData {
  report_name: string
  report_type: string
  period_type: string
  period_start: string
  period_end: string
}

export interface ReportListParams {
  report_type?: string
  period_type?: string
  page?: number
  size?: number
}