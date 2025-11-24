import { Navigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

function RoleDecider() {

    const token = localStorage.getItem("token");
    const decode = jwtDecode(token)
    const role = decode.role;
    
    if(!token) return <Navigate to="/login" />;
    if (role === "admin") return <Navigate to="/admin" />;
    if (role === "user") return <Navigate to="/user" />;
    return <Navigate to="/login" />;

}

export default RoleDecider
