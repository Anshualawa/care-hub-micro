
import React from 'react';
import Layout from '@/components/Layout';
import UserManagement from '@/components/users/UserManagement';
import { Navigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';

export default function UserManagementPage() {
  const { currentUser, loading, hasRole } = useAuth();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!currentUser) {
    return <Navigate to="/login" replace />;
  }

  if (!hasRole(['admin', 'superadmin'])) {
    return <Navigate to="/unauthorized" replace />;
  }

  return (
    <Layout>
      <UserManagement />
    </Layout>
  );
}
