import React from "react";
import { Navigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

export const ProtectedRoutes = ({children,role})=>{
    const token = localStorage.getItem("token");

    if(!token) 
        return <Navigate to="/login" replace/>
    
    try {
        const decoded = jwtDecode(token)
        const userRole = decoded.Role || decoded.role;

        if(role && userRole!==role){
            return <Navigate to="/" replace />;
        }
        return children
    } catch (error) {
        return <Navigate to="/" />;
    }
}