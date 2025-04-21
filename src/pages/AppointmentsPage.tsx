
import AppointmentList from "@/components/appointments/AppointmentList";
import Layout from "@/components/Layout";
import { Navigate } from "react-router-dom";
import { useAuth } from "@/context/AuthContext";

export default function AppointmentsPage() {
  const { currentUser, loading } = useAuth();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!currentUser) {
    return <Navigate to="/login" replace />;
  }

  return (
    <Layout>
      <div className="space-y-6">
        <h1 className="text-3xl font-bold tracking-tight">Appointment Management</h1>
        <AppointmentList />
      </div>
    </Layout>
  );
}
