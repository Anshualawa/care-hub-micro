import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Index from "./pages/Index";
import NotFound from "./pages/NotFound";
import LoginPage from "./pages/LoginPage";
import SignUpPage from "./pages/SignUpPage";
import DashboardPage from "./pages/DashboardPage";
import PatientsPage from "./pages/PatientsPage";
import AppointmentsPage from "./pages/AppointmentsPage";
import UnauthorizedPage from "./pages/UnauthorizedPage";
import DoctorsPage from "./pages/DoctorsPage";
import DoctorPortfolioPage from "./pages/DoctorPortfolioPage";
import BlogsPage from "./pages/BlogsPage";
import BlogDetailPage from "./pages/BlogDetailPage";
import BlogFormPage from "./pages/BlogFormPage";
import { AuthProvider } from "./context/AuthContext";
import RoleBasedRoute from "./components/RoleBasedRoute";

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <AuthProvider>
      <TooltipProvider>
        <Toaster />
        <Sonner />
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/signup" element={<SignUpPage />} />
            <Route path="/unauthorized" element={<UnauthorizedPage />} />
            
            <Route path="/dashboard" element={<DashboardPage />} />
            
            <Route path="/patients" element={
              <RoleBasedRoute allowedRoles={['admin', 'superadmin', 'doctor', 'nurse', 'intern']}>
                <PatientsPage />
              </RoleBasedRoute>
            } />
            
            <Route path="/appointments" element={<AppointmentsPage />} />
            
            <Route path="/doctors" element={<DoctorsPage />} />
            <Route path="/doctors/:doctorId" element={<DoctorPortfolioPage />} />
            
            <Route path="/blogs" element={<BlogsPage />} />
            <Route path="/blogs/:blogId" element={<BlogDetailPage />} />
            
            <Route path="/blogs/new" element={
              <RoleBasedRoute allowedRoles={['admin', 'superadmin', 'doctor']}>
                <BlogFormPage />
              </RoleBasedRoute>
            } />
            <Route path="/blogs/edit/:blogId" element={
              <RoleBasedRoute allowedRoles={['admin', 'superadmin', 'doctor']}>
                <BlogFormPage />
              </RoleBasedRoute>
            } />
            
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </TooltipProvider>
    </AuthProvider>
  </QueryClientProvider>
);

export default App;
