
import React from 'react';
import Layout from '@/components/Layout';
import SystemSettings from '@/components/settings/SystemSettings';
import { Navigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';

export default function SystemSettingsPage() {
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
      <SystemSettings />
    </Layout>
  );
}
