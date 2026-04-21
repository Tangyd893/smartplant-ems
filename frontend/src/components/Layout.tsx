import { Outlet, NavLink } from 'react-router-dom'

export default function Layout() {
  return (
    <div className="app-layout">
      <aside className="sidebar">
        <div className="sidebar-logo">⚡ SmartPlant</div>
        <nav>
          <NavLink to="/dashboard" className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}>
            📊 数据看板
          </NavLink>
          <NavLink to="/devices" className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}>
            🏭 设备管理
          </NavLink>
          <NavLink to="/energy" className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}>
            ⚡ 能耗分析
          </NavLink>
          <NavLink to="/reports" className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}>
            📄 报表管理
          </NavLink>
        </nav>
      </aside>
      <main className="main-content">
        <header className="topbar">
          <span style={{ fontSize: '0.9rem', color: 'var(--text-muted)' }}>智慧工厂能源管理系统</span>
        </header>
        <div className="page-content">
          <Outlet />
        </div>
      </main>
    </div>
  )
}