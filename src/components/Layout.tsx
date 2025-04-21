
import { Outlet, Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import { Button } from "./ui/button";
import { useToast } from "./ui/use-toast";
import { User } from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const { currentUser, logout, hasRole } = useAuth();
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
          
          <nav className="hidden md:flex items-center gap-6 text-sm">
            <Link to="/" className="font-medium transition-colors hover:text-teal-600">
              Home
            </Link>
            
            <Link to="/doctors" className="font-medium transition-colors hover:text-teal-600">
              Doctors
            </Link>
            
            <Link to="/blogs" className="font-medium transition-colors hover:text-teal-600">
              Blog
            </Link>
            
            {currentUser && (
              <>
                <Link to="/dashboard" className="font-medium transition-colors hover:text-teal-600">
                  Dashboard
                </Link>
                
                {hasRole(['admin', 'superadmin', 'doctor', 'nurse', 'intern']) && (
                  <Link to="/patients" className="font-medium transition-colors hover:text-teal-600">
                    Patients
                  </Link>
                )}
                
                <Link to="/appointments" className="font-medium transition-colors hover:text-teal-600">
                  Appointments
                </Link>
                
                {hasRole(['admin', 'superadmin']) && (
                  <Link to="/users" className="font-medium transition-colors hover:text-teal-600">
                    Users
                  </Link>
                )}
                
                {hasRole('superadmin') && (
                  <Link to="/settings" className="font-medium transition-colors hover:text-teal-600">
                    System Settings
                  </Link>
                )}
              </>
            )}
          </nav>
          
          <div className="flex items-center gap-4">
            {currentUser ? (
              <div className="flex items-center gap-4">
                <div className="hidden md:flex items-center gap-2">
                  <Avatar className="h-8 w-8">
                    <AvatarFallback className="bg-primary/10 text-primary">
                      {currentUser.name?.split(' ').map(n => n[0]).join('')}
                    </AvatarFallback>
                  </Avatar>
                  <div>
                    <div className="text-sm font-medium">{currentUser.name}</div>
                    <div className="text-xs text-gray-500">{currentUser.role}</div>
                  </div>
                </div>
                <Button onClick={handleLogout} variant="ghost" size="sm">
                  Logout
                </Button>
              </div>
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
        <div className="container mx-auto py-8 px-4 sm:px-6">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div>
              <h3 className="text-lg font-bold mb-4">CareHub Hospital</h3>
              <p className="text-sm text-gray-500">
                Providing exceptional healthcare services with compassion and excellence.
              </p>
            </div>
            
            <div>
              <h3 className="text-lg font-bold mb-4">Quick Links</h3>
              <ul className="space-y-2 text-sm">
                <li>
                  <Link to="/" className="text-gray-500 hover:text-primary">
                    Home
                  </Link>
                </li>
                <li>
                  <Link to="/doctors" className="text-gray-500 hover:text-primary">
                    Our Doctors
                  </Link>
                </li>
                <li>
                  <Link to="/blogs" className="text-gray-500 hover:text-primary">
                    Blog
                  </Link>
                </li>
                {currentUser && (
                  <li>
                    <Link to="/appointments" className="text-gray-500 hover:text-primary">
                      Appointments
                    </Link>
                  </li>
                )}
              </ul>
            </div>
            
            <div>
              <h3 className="text-lg font-bold mb-4">Contact</h3>
              <address className="not-italic text-sm text-gray-500">
                <p>123 Healthcare Avenue</p>
                <p>Medical District, MD 12345</p>
                <p className="mt-2">Email: contact@carehub.com</p>
                <p>Phone: (555) 123-4567</p>
              </address>
            </div>
          </div>
          
          <div className="mt-8 pt-8 border-t text-center text-sm text-gray-500">
            © 2025 CareHub Hospital. All rights reserved.
          </div>
        </div>
      </footer>
    </div>
  );
}
