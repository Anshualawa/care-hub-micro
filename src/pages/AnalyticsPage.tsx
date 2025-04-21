
import React from 'react';
import Layout from '@/components/Layout';
import AnalyticsDashboard from '@/components/analytics/AnalyticsDashboard';
import { Navigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';

export default function AnalyticsPage() {
  const { currentUser, loading, hasRole } = useAuth();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!currentUser) {
    return <Navigate to="/login" replace />;
  }

  if (!hasRole(['admin', 'doctor', 'superadmin'])) {
    return <Navigate to="/unauthorized" replace />;
  }

  return (
    <Layout>
      <AnalyticsDashboard />
    </Layout>
  );
}
