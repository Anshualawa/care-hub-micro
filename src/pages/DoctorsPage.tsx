
import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { doctorService } from '../services/api';
import Layout from '@/components/Layout';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { useToast } from '@/components/ui/use-toast';
import { User } from 'lucide-react';

export default function DoctorsPage() {
  const { toast } = useToast();
  
  const { data: doctors, isLoading, error } = useQuery({
    queryKey: ['doctors'],
    queryFn: doctorService.getAllDoctors,
    retry: false, // Prevent multiple retries that could cause render loops
    onError: (err) => {
      // Handle error in onError callback instead of during render
      console.error('Failed to load doctors:', err);
    }
  });

  // Use useEffect to show toast on error
  useEffect(() => {
    if (error) {
      toast({
        variant: "destructive",
        title: "Error",
        description: "Failed to load doctors. API may not be available.",
      });
    }
  }, [error, toast]);

  if (isLoading) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading doctors...</div>
        </div>
      </Layout>
    );
  }

  // Handle error state without calling toast during render
  if (error) {
    return (
      <Layout>
        <div className="text-center py-10">
          <h2 className="text-xl font-bold">Error loading doctors</h2>
          <p className="text-muted-foreground">The doctors API endpoint may not be available.</p>
          <div className="mt-4">
            <Button variant="outline" onClick={() => window.location.reload()}>
              Try Again
            </Button>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="container mx-auto py-8">
        <div className="text-center mb-10">
          <h1 className="text-3xl font-bold mb-2">Our Medical Team</h1>
          <p className="text-muted-foreground max-w-2xl mx-auto">
            Meet our highly qualified and experienced medical professionals dedicated to 
            providing the best healthcare services at CareHub Hospital.
          </p>
        </div>
        
        {doctors && doctors.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {doctors.map((doctor: any) => (
              <Card key={doctor.id} className="overflow-hidden">
                <div className="aspect-square relative bg-muted">
                  {doctor.profileImage ? (
                    <img 
                      src={doctor.profileImage} 
                      alt={doctor.name}
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <div className="flex items-center justify-center h-full bg-primary/10">
                      <User className="h-24 w-24 text-primary/30" />
                    </div>
                  )}
                </div>
                
                <CardHeader className="pb-2">
                  <CardTitle>{doctor.name}</CardTitle>
                  <CardDescription>{doctor.specialization || doctor.department}</CardDescription>
                </CardHeader>
                
                <CardContent>
                  {doctor.department && (
                    <Badge className="mb-2">{doctor.department}</Badge>
                  )}
                  {doctor.bio && (
                    <p className="text-sm text-muted-foreground line-clamp-2 mt-2">
                      {doctor.bio}
                    </p>
                  )}
                </CardContent>
                
                <CardFooter>
                  <Link to={`/doctors/${doctor.id}`} className="w-full">
                    <Button variant="outline" className="w-full">
                      View Profile
                    </Button>
                  </Link>
                </CardFooter>
              </Card>
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            <h2 className="text-xl font-semibold mb-2">No doctors found</h2>
            <p className="text-muted-foreground">
              Our doctor directory is currently being updated or the API may not be available.
            </p>
          </div>
        )}
      </div>
    </Layout>
  );
}
