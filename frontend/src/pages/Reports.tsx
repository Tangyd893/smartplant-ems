import { useState, useEffect } from 'react'
import axios from 'axios'

interface Report {
  id: number
  report_name: string
  report_type: string
  period_type: string
  period_start: string
  period_end: string
  status: number
}

const REPORT_TYPE_MAP: Record<string, string> = {
  daily: '日报表',
  monthly: '月报表',
  yearly: '年报表',
}

export default function Reports() {
  const [reports, setReports] = useState<Report[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    axios.get('/api/report/list?page=1&size=50')
      .then(res => {
        setReports(res.data?.data?.records || [])
        setLoading(false)
      })
      .catch(() => setLoading(false))
  }, [])

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem' }}>
        <h1 className="page-title" style={{ margin: 0 }}>报表管理</h1>
        <button className="btn btn-primary">+ 生成报表</button>
      </div>

      <div className="card">
        {loading ? (
          <div style={{ padding: '2rem', textAlign: 'center', color: 'var(--text-muted)' }}>加载中...</div>
        ) : reports.length === 0 ? (
          <div style={{ padding: '2rem', textAlign: 'center', color: 'var(--text-muted)' }}>暂无报表数据</div>
        ) : (
          <table>
            <thead>
              <tr>
                <th>报表名称</th>
                <th>类型</th>
                <th>周期</th>
                <th>时间范围</th>
                <th>状态</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {reports.map(r => (
                <tr key={r.id}>
                  <td>{r.report_name}</td>
                  <td>{REPORT_TYPE_MAP[r.report_type] || r.report_type}</td>
                  <td>{r.period_type}</td>
                  <td style={{ fontSize: '0.85rem', color: 'var(--text-muted)' }}>
                    {r.period_start?.split('T')[0]} ~ {r.period_end?.split('T')[0]}
                  </td>
                  <td>
                    <span className={`badge ${r.status === 1 ? 'badge-green' : 'badge-yellow'}`}>
                      {r.status === 1 ? '已完成' : '生成中'}
                    </span>
                  </td>
                  <td>
                    <button className="btn btn-secondary" style={{ fontSize: '0.8rem', padding: '0.25rem 0.75rem', marginRight: '0.5rem' }}>
                      下载
                    </button>
                    <button className="btn btn-secondary" style={{ fontSize: '0.8rem', padding: '0.25rem 0.75rem', color: 'var(--danger)' }}>
                      删除
                    </button>
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