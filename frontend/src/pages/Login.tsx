import { useState } from 'react'

export default function Login() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (username === 'admin' && password === 'admin123') {
      localStorage.setItem('token', 'mock-token-123')
      window.location.href = '/dashboard'
    } else {
      alert('用户名或密码错误')
    }
  }

  return (
    <div style={{ minHeight: '100vh', display: 'flex', alignItems: 'center', justifyContent: 'center', background: 'var(--bg-primary)' }}>
      <div className="card" style={{ width: '360px', padding: '2rem' }}>
        <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
          <div style={{ fontSize: '2rem', marginBottom: '0.5rem' }}>⚡</div>
          <h1 style={{ fontSize: '1.25rem', fontWeight: 600 }}>SmartPlant EMS</h1>
          <p style={{ fontSize: '0.85rem', color: 'var(--text-muted)' }}>智慧工厂能源管理系统</p>
        </div>
        <form onSubmit={handleSubmit}>
          <div style={{ marginBottom: '1rem' }}>
            <label style={{ display: 'block', fontSize: '0.85rem', marginBottom: '0.375rem', color: 'var(--text-secondary)' }}>用户名</label>
            <input
              type="text"
              value={username}
              onChange={e => setUsername(e.target.value)}
              style={{ width: '100%', padding: '0.5rem 0.75rem', border: '1px solid var(--border)', borderRadius: '6px', fontSize: '0.9rem' }}
            />
          </div>
          <div style={{ marginBottom: '1.5rem' }}>
            <label style={{ display: 'block', fontSize: '0.85rem', marginBottom: '0.375rem', color: 'var(--text-secondary)' }}>密码</label>
            <input
              type="password"
              value={password}
              onChange={e => setPassword(e.target.value)}
              style={{ width: '100%', padding: '0.5rem 0.75rem', border: '1px solid var(--border)', borderRadius: '6px', fontSize: '0.9rem' }}
            />
          </div>
          <button type="submit" className="btn btn-primary" style={{ width: '100%' }}>
            登录
          </button>
        </form>
        <p style={{ marginTop: '1rem', fontSize: '0.75rem', color: 'var(--text-muted)', textAlign: 'center' }}>
          测试账号: admin / admin123
        </p>
      </div>
    </div>
  )
}