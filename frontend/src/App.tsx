import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import Layout from './components/Layout'
import Dashboard from './pages/Dashboard'
import DeviceList from './pages/DeviceList'
import DeviceDetail from './pages/DeviceDetail'
import EnergyStats from './pages/EnergyStats'
import Reports from './pages/Reports'
import Login from './pages/Login'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<Layout />}>
          <Route index element={<Navigate to="/dashboard" replace />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="devices" element={<DeviceList />} />
          <Route path="devices/:id" element={<DeviceDetail />} />
          <Route path="energy" element={<EnergyStats />} />
          <Route path="reports" element={<Reports />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}