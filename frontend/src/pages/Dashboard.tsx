import { useState, useEffect } from 'react'

const mockStats = {
  totalDevices: 12,
  onlineDevices: 10,
  todayEnergy: 4523.8,
  todayCost: 2488.09,
  alerts: 3,
  avgPower: 125.3,
}

export default function Dashboard() {
  const [time, setTime] = useState(new Date())

  useEffect(() => {
    const timer = setInterval(() => setTime(new Date()), 1000)
    return () => clearInterval(timer)
  }, [])

  return (
    <div>
      <div className="card" style={{ marginBottom: '1rem', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <div>
          <span style={{ fontSize: '0.85rem', color: 'var(--text-muted)' }}>实时监控</span>
          <h2 style={{ fontSize: '1rem', fontWeight: 600 }}>{time.toLocaleString('zh-CN')}</h2>
        </div>
        <div className="badge badge-green">系统运行中</div>
      </div>

      <div className="stat-grid">
        <div className="stat-card">
          <div className="stat-label">设备总数</div>
          <div className="stat-value">{mockStats.totalDevices}</div>
          <div style={{ fontSize: '0.75rem', color: 'var(--text-muted)' }}>在线 {mockStats.onlineDevices}</div>
        </div>
        <div className="stat-card">
          <div className="stat-label">今日能耗</div>
          <div className="stat-value">{mockStats.todayEnergy.toFixed(1)}<span className="stat-unit">kWh</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-label">今日成本</div>
          <div className="stat-value">¥{mockStats.todayCost.toFixed(2)}</div>
        </div>
        <div className="stat-card">
          <div className="stat-label">活跃告警</div>
          <div className="stat-value" style={{ color: mockStats.alerts > 0 ? 'var(--danger)' : 'inherit' }}>{mockStats.alerts}</div>
        </div>
        <div className="stat-card">
          <div className="stat-label">平均功率</div>
          <div className="stat-value">{mockStats.avgPower.toFixed(1)}<span className="stat-unit">kW</span></div>
        </div>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
        <div className="card">
          <div style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '0.75rem' }}>🏭 设备状态概览</div>
          <table>
            <thead>
              <tr>
                <th>设备</th>
                <th>类型</th>
                <th>状态</th>
              </tr>
            </thead>
            <tbody>
              <tr><td>空压机A1</td><td>空压机</td><td><span className="badge badge-green">运行中</span></td></tr>
              <tr><td>注塑机B1</td><td>注塑机</td><td><span className="badge badge-green">运行中</span></td></tr>
              <tr><td>激光切割机C1</td><td>激光切割机</td><td><span className="badge badge-yellow">维护中</span></td></tr>
              <tr><td>空压机A2</td><td>空压机</td><td><span className="badge badge-green">运行中</span></td></tr>
            </tbody>
          </table>
        </div>

        <div className="card">
          <div style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '0.75rem' }}>⚡ 实时功率TOP5</div>
          <table>
            <thead>
              <tr>
                <th>设备</th>
                <th>功率(kW)</th>
              </tr>
            </thead>
            <tbody>
              <tr><td>空压机A1</td><td style={{ fontWeight: 600 }}>150.0</td></tr>
              <tr><td>空压机A2</td><td style={{ fontWeight: 600 }}>148.5</td></tr>
              <tr><td>注塑机B1</td><td style={{ fontWeight: 600 }}>75.2</td></tr>
              <tr><td>注塑机B2</td><td style={{ fontWeight: 600 }}>72.8</td></tr>
              <tr><td>激光切割机C1</td><td style={{ fontWeight: 600 }}>45.3</td></tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}