import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
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

const DEVICE_TYPE_MAP: Record<string, string> = {
  air_compressor: '空压机',
  injection_machine: '注塑机',
  laser: '激光切割机',
}

const STATUS_MAP: Record<number, { label: string; class: string }> = {
  1: { label: '运行中', class: 'badge-green' },
  2: { label: '维护中', class: 'badge-yellow' },
  3: { label: '故障', class: 'badge-red' },
  0: { label: '停用', class: 'badge-blue' },
}

export default function DeviceList() {
  const [devices, setDevices] = useState<Device[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    axios.get('/api/device/list?page=1&size=50')
      .then(res => {
        setDevices(res.data?.data?.records || [])
        setLoading(false)
      })
      .catch(() => setLoading(false))
  }, [])

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem' }}>
        <h1 className="page-title" style={{ margin: 0 }}>设备管理</h1>
        <button className="btn btn-primary">+ 添加设备</button>
      </div>

      <div className="card">
        {loading ? (
          <div style={{ padding: '2rem', textAlign: 'center', color: 'var(--text-muted)' }}>加载中...</div>
        ) : devices.length === 0 ? (
          <div style={{ padding: '2rem', textAlign: 'center', color: 'var(--text-muted)' }}>暂无设备数据</div>
        ) : (
          <table>
            <thead>
              <tr>
                <th>设备编号</th>
                <th>设备名称</th>
                <th>类型</th>
                <th>位置</th>
                <th>额定功率(kW)</th>
                <th>状态</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {devices.map(d => (
                <tr key={d.id}>
                  <td style={{ fontFamily: 'monospace' }}>{d.device_code}</td>
                  <td>{d.device_name}</td>
                  <td>{DEVICE_TYPE_MAP[d.device_type] || d.device_type}</td>
                  <td>{d.location}</td>
                  <td>{d.power_rating}</td>
                  <td><span className={`badge ${STATUS_MAP[d.status]?.class}`}>{STATUS_MAP[d.status]?.label}</span></td>
                  <td>
                    <Link to={`/devices/${d.id}`} className="btn btn-secondary" style={{ fontSize: '0.8rem', padding: '0.25rem 0.75rem' }}>详情</Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  )
}