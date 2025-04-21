
import { Outlet, Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import { Button } from "./ui/button";
import { useToast } from "./ui/use-toast";

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const { currentUser, logout } = useAuth();
  const navigate = useNavigate();
  const { toast } = useToast();

  const handleLogout = () => {
    logout();
    toast({
      title: "Logged out",
      description: "You have been successfully logged out",
    });
    navigate("/");
  };

  return (
    <div className="flex min-h-screen flex-col">
      <header className="sticky top-0 z-10 bg-white border-b shadow-sm">
        <div className="flex h-16 items-center justify-between px-4 sm:px-6">
          <div className="flex items-center">
            <Link to="/" className="flex items-center">
              <span className="text-xl font-bold text-teal-600">CareHub</span>
            </Link>
          </div>
          
          {currentUser && (
            <nav className="hidden md:flex items-center gap-6 text-sm">
              <Link to="/dashboard" className="font-medium transition-colors hover:text-teal-600">
                Dashboard
              </Link>
              <Link to="/patients" className="font-medium transition-colors hover:text-teal-600">
                Patients
              </Link>
              <Link to="/appointments" className="font-medium transition-colors hover:text-teal-600">
                Appointments
              </Link>
            </nav>
          )}
          
          <div className="flex items-center gap-4">
            {currentUser ? (
              <Button onClick={handleLogout} variant="ghost" size="sm">
                Logout
              </Button>
            ) : (
              <Link to="/login">
                <Button variant="outline" size="sm">
                  Login
                </Button>
              </Link>
            )}
          </div>
        </div>
      </header>
      
      <main className="flex-1">
        <div className="container mx-auto p-4 sm:p-6">
          {children}
        </div>
      </main>
      
      <footer className="border-t bg-gray-50">
        <div className="container mx-auto py-4 px-4 sm:px-6 text-center text-sm text-gray-500">
          © 2025 CareHub Microservice. All rights reserved.
        </div>
      </footer>
    </div>
  );
}
