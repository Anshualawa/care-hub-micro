
import PatientList from "@/components/patients/PatientList";
import Layout from "@/components/Layout";
import { Navigate } from "react-router-dom";
import { useAuth } from "@/context/AuthContext";

export default function PatientsPage() {
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
        <h1 className="text-3xl font-bold tracking-tight">Patient Management</h1>
        <PatientList />
      </div>
    </Layout>
  );
}
