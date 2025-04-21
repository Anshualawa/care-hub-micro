
import Dashboard from "@/components/Dashboard";
import Layout from "@/components/Layout";
import { Navigate } from "react-router-dom";
import { useAuth } from "@/context/AuthContext";

export default function DashboardPage() {
  const { currentUser, loading } = useAuth();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!currentUser) {
    return <Navigate to="/login" replace />;
  }

  return (
    <Layout>
      <Dashboard />
    </Layout>
  );
}
