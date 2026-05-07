export interface EnergyRecord {
  id: number
  device_id: number
  device_name?: string
  record_time: string
  power_kw: number
  energy_kwh: number
  voltage_v: number
  current_a: number
  power_factor: number
  created_at?: string
}

export interface EnergyStats {
  device_id: number
  device_name: string
  total_energy_kwh: number
  avg_power_kw: number
  max_power_kw: number
  min_power_kw: number
  record_count: number
}

export interface EnergyListParams {
  device_id?: number
  start_time?: string
  end_time?: string
  page?: number
  size?: number
}