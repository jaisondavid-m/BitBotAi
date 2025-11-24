import React from 'react'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from '../pages/Login';
import Register from '../pages/Register';
import AdminHome from '../pages/AdminHome';
import StudentHome from '../pages/StudentHome';
import { ProtectedRoutes } from '../components/ProtectedRoutes';
import RoleDecider from '../components/RoleDecider';

function App() {
  return (
    <div>
        <Router>
            <Routes>
                <Route path="/login" element={<Login/>} />
                <Route path="/register" element={<Register />} />
                <Route path='/admin' element={
                  <ProtectedRoutes role="admin">
                    <AdminHome/>
                  </ProtectedRoutes>
                }/>
                <Route path='/user' element={
                  <ProtectedRoutes role="user">
                    <StudentHome/>
                  </ProtectedRoutes>
                }/>
                <Route
                  path='/'
                  element={
                    <ProtectedRoutes>
                      <RoleDecider/>
                    </ProtectedRoutes>
                  }
                />
            </Routes>
        </Router>
    </div>
  )
}

export default App
