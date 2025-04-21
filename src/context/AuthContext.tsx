
import React, { createContext, useState, useEffect, useContext } from 'react';
import { authService } from '../services/api';

interface User {
  id: number;
  name: string;
  role: string;
  email?: string;
  phone?: string;
  department?: string;
  specialization?: string;
  bio?: string;
  education?: string[];
  experience?: string[];
  profileImage?: string;
}

interface AuthContextType {
  currentUser: User | null;
  loading: boolean;
  login: (username: string, password: string) => Promise<any>;
  logout: () => void;
  isAuthenticated: boolean;
  hasRole: (roles: string | string[]) => boolean;
}

const AuthContext = createContext<AuthContextType>({
  currentUser: null,
  loading: true,
  login: async () => ({}),
  logout: () => {},
  isAuthenticated: false,
  hasRole: () => false,
});

export const useAuth = () => useContext(AuthContext);

export const AuthProvider: React.FC<{children: React.ReactNode}> = ({ children }) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const user = authService.getCurrentUser();
    setCurrentUser(user);
    setLoading(false);
  }, []);

  const login = async (username: string, password: string) => {
    try {
      const data = await authService.login(username, password);
      setCurrentUser(data.user);
      return data;
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    authService.logout();
    setCurrentUser(null);
  };

  // Helper function to check if user has specific role(s)
  const hasRole = (roles: string | string[]) => {
    if (!currentUser) return false;
    
    if (typeof roles === 'string') {
      return currentUser.role === roles;
    }
    
    return roles.includes(currentUser.role);
  };

  const value = {
    currentUser,
    loading,
    login,
    logout,
    isAuthenticated: !!currentUser,
    hasRole,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
