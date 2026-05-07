import { useState, useEffect } from 'react'
import axios from 'axios'

interface Stats {
  total_energy_kwh: number
  avg_power_kw: number
  max_power_kw: number
  min_power_kw: number
  total_cost: number
  peak_hours: number
  off_peak_hours: number
}

export default function EnergyStats() {
  const [stats, setStats] = useState<Stats | null>(null)
  const [period, setPeriod] = useState('day')

  useEffect(() => {
    axios.get(`/api/energy/stats?period=${period}`).then(res => setStats(res.data?.data))
  }, [period])

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem' }}>
        <h1 className="page-title" style={{ margin: 0 }}>能耗分析</h1>
        <select className="btn btn-secondary" value={period} onChange={e => setPeriod(e.target.value)}>
          <option value="day">今日</option>
          <option value="month">本月</option>
          <option value="year">本年</option>
        </select>
      </div>

      <div className="stat-grid">
        <div className="stat-card">
          <div className="stat-label">总能耗</div>
          <div className="stat-value">{stats?.total_energy_kwh?.toFixed(1) || '—'}<span className="stat-unit">kWh</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-label">总成本</div>
          <div className="stat-value">¥{stats?.total_cost?.toFixed(2) || '—'}</div>
        </div>
        <div className="stat-card">
          <div className="stat-label">平均功率</div>
          <div className="stat-value">{stats?.avg_power_kw?.toFixed(1) || '—'}<span className="stat-unit">kW</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-label">最大功率</div>
          <div className="stat-value">{stats?.max_power_kw?.toFixed(1) || '—'}<span className="stat-unit">kW</span></div>
        </div>
      </div>

      <div className="card">
        <h2 style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '0.75rem' }}>📈 能耗趋势</h2>
        <div style={{ height: '200px', display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'var(--text-muted)', border: '1px dashed var(--border)', borderRadius: '8px' }}>
          图表区域（ECharts 接入后展示）
        </div>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', marginTop: '1rem' }}>
        <div className="card">
          <h2 style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '0.75rem' }}>⏰ 时段分布</h2>
          <table>
            <thead><tr><th>时段</th><th>小时数</th></tr></thead>
            <tbody>
              <tr><td>峰时段</td><td style={{ fontWeight: 600 }}>{stats?.peak_hours || 0}h</td></tr>
              <tr><td>谷时段</td><td style={{ fontWeight: 600 }}>{stats?.off_peak_hours || 0}h</td></tr>
            </tbody>
          </table>
        </div>
        <div className="card">
          <h2 style={{ fontSize: '0.9rem', fontWeight: 500, marginBottom: '0.75rem' }}>⚡ 功率区间</h2>
          <table>
            <thead><tr><th>指标</th><th>值(kW)</th></tr></thead>
            <tbody>
              <tr><td>最大功率</td><td style={{ fontWeight: 600 }}>{stats?.max_power_kw || 0}</td></tr>
              <tr><td>最小功率</td><td style={{ fontWeight: 600 }}>{stats?.min_power_kw || 0}</td></tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}