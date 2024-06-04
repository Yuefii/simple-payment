import React, { ReactNode } from "react";
import { Navigate } from "react-router-dom";

interface ProtectedRouteProps {
  children: ReactNode;
}

const Auth: React.FC<ProtectedRouteProps> = ({ children }) => {
  const token = localStorage.getItem("token");

  if (!token) {
    console.warn("Unauthorized.");
    return <Navigate to="/auth/login" />;
  }

  return <>{children}</>;
};

export default Auth;
