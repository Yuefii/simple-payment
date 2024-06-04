import { Route, Routes } from 'react-router-dom'

import Home from './pages/Home'
import Login from './pages/auth/Login'
import Register from './pages/auth/Register'

export default function Router() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/auth/register" element={<Register />} />
      <Route path="/auth/login" element={<Login />} />
    </Routes>
  )
}