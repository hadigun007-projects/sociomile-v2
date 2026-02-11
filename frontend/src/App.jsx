import { Routes, Route } from 'react-router-dom'
import Login from './pages/Login'

function Home() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-4xl font-bold text-blue-600 mb-4">Welcome to React + Tailwind</h1>
      <p className="text-gray-700">This is a basic setup with React Router, Axios, and TanStack Query.</p>
      <a href="/login" className="mt-4 text-blue-500 hover:underline">Go to Login</a>
    </div>
  )
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
    </Routes>
  )
}

export default App
