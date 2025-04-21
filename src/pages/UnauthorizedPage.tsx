
import React from 'react';
import { Link } from 'react-router-dom';
import Layout from '@/components/Layout';
import { Button } from '@/components/ui/button';
import { useAuth } from '../context/AuthContext';
import { ShieldAlert } from 'lucide-react';

export default function UnauthorizedPage() {
  const { currentUser } = useAuth();
  
  return (
    <Layout>
      <div className="flex flex-col items-center justify-center py-12 px-4 text-center">
        <div className="rounded-full bg-red-100 p-4 mb-6">
          <ShieldAlert className="text-red-600 h-10 w-10" />
        </div>
        <h1 className="text-3xl font-bold mb-2">Access Denied</h1>
        <p className="mb-6 text-gray-600">
          Sorry, you don't have permission to access this page.
        </p>
        <p className="mb-6 text-gray-600">
          Your current role is: <span className="font-medium">{currentUser?.role}</span>
        </p>
        <div className="space-x-4">
          <Button onClick={() => window.history.back()} variant="outline">
            Go Back
          </Button>
          <Link to="/dashboard">
            <Button>Go to Dashboard</Button>
          </Link>
        </div>
      </div>
    </Layout>
  );
}
