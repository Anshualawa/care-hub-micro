
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

interface RoleBasedRouteProps {
  children: React.ReactNode;
  allowedRoles: string[];
  redirectPath?: string;
  requireAuth?: boolean;
}

export default function RoleBasedRoute({ 
  children, 
  allowedRoles, 
  redirectPath = '/unauthorized',
  requireAuth = true
}: RoleBasedRouteProps) {
  const { isAuthenticated, hasRole, loading } = useAuth();

  if (loading) {
    return <div className="flex justify-center p-8">Loading...</div>;
  }

  // If authentication is required and user is not authenticated
  if (requireAuth && !isAuthenticated) {
    return <Navigate to="/login" />;
  }

  // If user doesn't have the required role
  if (requireAuth && !hasRole(allowedRoles)) {
    return <Navigate to={redirectPath} />;
  }

  return <>{children}</>;
}
