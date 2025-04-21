
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

interface RoleBasedRouteProps {
  children: React.ReactNode;
  allowedRoles: string[];
  redirectPath?: string;
}

export default function RoleBasedRoute({ 
  children, 
  allowedRoles, 
  redirectPath = '/unauthorized'
}: RoleBasedRouteProps) {
  const { isAuthenticated, hasRole, loading } = useAuth();

  if (loading) {
    return <div className="flex justify-center p-8">Loading...</div>;
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }

  if (!hasRole(allowedRoles)) {
    return <Navigate to={redirectPath} />;
  }

  return <>{children}</>;
}
