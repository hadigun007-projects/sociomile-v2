import { Routes, Route } from 'react-router-dom'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import ProtectedRoute from './components/ProtectedRoute'
import DashboardLayout from './layouts/DashboardLayout'
import Users from './pages/Users'
import Conversations from './pages/Conversations'
import Tenants from './pages/Tenants'
import Tickets from './pages/Tickets'
import Customers from './pages/Customers'
import ChannelSimulator from './pages/ChannelSimulator'

function Home() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-4xl font-bold text-blue-600 mb-4">Welcome to React + Tailwind</h1>
      <p className="text-gray-700 mb-8">This is a basic setup with React Router, Axios, and TanStack Query.</p>
      <div className="flex gap-4">
        <a href="/login" className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">Login</a>
        <a href="/dashboard" className="px-6 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors">Dashboard</a>
      </div>
    </div>
  )
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/customer-simulator" element={<ChannelSimulator />} />

      {/* Protected Routes */}
      <Route element={<ProtectedRoute />}>
        {/* Dashboard Layout wraps all dashboard child routes */}
        <Route path="/dashboard" element={<DashboardLayout />}>
          <Route index element={<Dashboard />} />
          <Route path="users" element={<Users />} />
          <Route path="conversations" element={<Conversations />} />
          <Route path="tenants" element={<Tenants />} />
          <Route path="tickets" element={<Tickets />} />
          <Route path="customers" element={<Customers />} />
        </Route>
      </Route>
    </Routes>
  )
}

export default App
