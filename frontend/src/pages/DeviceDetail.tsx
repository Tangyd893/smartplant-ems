import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import axios from 'axios'

interface Device {
  id: number
  device_code: string
  device_name: string
  device_type: string
  location: string
  status: number
  power_rating: number
}

interface RealtimeData {
  device_id: number
  power_kw: number
  voltage_v: number
  current_a: number
  power_factor: number
  timestamp: string
}

export default function DeviceDetail() {
  const { id } = useParams()
  const [device, setDevice] = useState<Device | null>(null)
  const [realtime, setRealtime] = useState<RealtimeData | null>(null)

  useEffect(() => {
    if (!id) return
    axios.get(`/api/device/${id}`).then(res => setDevice(res.data?.data))
    axios.get(`/api/device/${id}/realtime`).then(res => setRealtime(res.data?.data))
  }, [id])

  if (!device) return <div style={{ padding: '2rem' }}>加载中...</div>

  return (
    <div>
      <h1 className="page-title">设备详情</h1>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
        <div className="card">
          <h2 style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '1rem' }}>基本信息</h2>
          <table>
            <tbody>
              <tr><td style={{ color: 'var(--text-muted)', width: '40%' }}>设备编号</td><td style={{ fontFamily: 'monospace' }}>{device.device_code}</td></tr>
              <tr><td style={{ color: 'var(--text-muted)' }}>设备名称</td><td>{device.device_name}</td></tr>
              <tr><td style={{ color: 'var(--text-muted)' }}>设备类型</td><td>{device.device_type}</td></tr>
              <tr><td style={{ color: 'var(--text-muted)' }}>安装位置</td><td>{device.location}</td></tr>
              <tr><td style={{ color: 'var(--text-muted)' }}>额定功率</td><td>{device.power_rating} kW</td></tr>
            </tbody>
          </table>
        </div>

        <div className="card">
          <h2 style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '1rem' }}>实时数据</h2>
          {realtime ? (
            <table>
              <tbody>
                <tr><td style={{ color: 'var(--text-muted)' }}>瞬时功率</td><td style={{ fontWeight: 600 }}>{realtime.power_kw} kW</td></tr>
                <tr><td style={{ color: 'var(--text-muted)' }}>电压</td><td>{realtime.voltage_v} V</td></tr>
                <tr><td style={{ color: 'var(--text-muted)' }}>电流</td><td>{realtime.current_a} A</td></tr>
                <tr><td style={{ color: 'var(--text-muted)' }}>功率因数</td><td>{realtime.power_factor}</td></tr>
                <tr><td style={{ color: 'var(--text-muted)' }}>采集时间</td><td>{realtime.timestamp}</td></tr>
              </tbody>
            </table>
          ) : <div style={{ color: 'var(--text-muted)' }}>加载中...</div>}
        </div>
      </div>
    </div>
  )
}