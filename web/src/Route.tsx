import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Register from './pages/Register'

export default function Router() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/auth/register" element={<Register />} />
    </Routes>
  )
}