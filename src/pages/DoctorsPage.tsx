
import React, { useEffect, useState } from 'react';
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

// Mock data to use when API is unavailable
const MOCK_DOCTORS = [
  {
    id: 1,
    name: "Dr. Jane Smith",
    specialization: "Cardiology",
    department: "Cardiology",
    bio: "Dr. Smith is a board-certified cardiologist with over 15 years of experience in treating heart conditions.",
    email: "jane.smith@carehub.com",
    phone: "555-123-4567",
    education: ["MD, Harvard Medical School", "Residency, Mayo Clinic"],
    experience: ["Senior Cardiologist, Mayo Clinic (2015-2020)", "Chief of Cardiology, CareHub Hospital (2020-Present)"],
  },
  {
    id: 2,
    name: "Dr. Robert Chen",
    specialization: "Neurology",
    department: "Neurology",
    bio: "Dr. Chen specializes in neurological disorders and has pioneered several treatment approaches for chronic conditions.",
    email: "robert.chen@carehub.com",
    phone: "555-234-5678",
    education: ["MD, Johns Hopkins University", "Fellowship, Cleveland Clinic"],
    experience: ["Neurologist, Cleveland Clinic (2013-2018)", "Senior Neurologist, CareHub Hospital (2018-Present)"],
    profileImage: "https://randomuser.me/api/portraits/men/32.jpg"
  },
  {
    id: 3,
    name: "Dr. Maria Rodriguez",
    specialization: "Pediatrics",
    department: "Pediatrics",
    bio: "Dr. Rodriguez has dedicated her career to children's health and is beloved by her young patients for her gentle approach.",
    email: "maria.rodriguez@carehub.com",
    phone: "555-345-6789",
    education: ["MD, Stanford University", "Residency, Children's Hospital of Philadelphia"],
    experience: ["Pediatrician, Boston Children's Hospital (2016-2021)", "Lead Pediatrician, CareHub Hospital (2021-Present)"],
    profileImage: "https://randomuser.me/api/portraits/women/45.jpg"
  },
  {
    id: 4,
    name: "Dr. James Wilson",
    specialization: "Orthopedic Surgery",
    department: "Orthopedics",
    bio: "Dr. Wilson is an orthopedic surgeon specializing in sports injuries and joint replacements.",
    email: "james.wilson@carehub.com",
    phone: "555-456-7890",
    education: ["MD, University of Michigan", "Orthopedic Fellowship, Hospital for Special Surgery"],
    experience: ["Orthopedic Surgeon, UCSF Medical Center (2012-2019)", "Chief of Orthopedics, CareHub Hospital (2019-Present)"],
  }
];

export default function DoctorsPage() {
  const { toast } = useToast();
  const [useMockData, setUseMockData] = useState(false);
  
  const { data: apiDoctors, isLoading, error } = useQuery({
    queryKey: ['doctors'],
    queryFn: doctorService.getAllDoctors,
    retry: 1, // Only retry once
    enabled: !useMockData, // Only run query if we're not using mock data
  });

  // Use mock data if API fails
  const doctors = useMockData ? MOCK_DOCTORS : apiDoctors;

  // Handle API error and switch to mock data
  useEffect(() => {
    if (error) {
      console.error('Failed to load doctors:', error);
      toast({
        variant: "destructive",
        title: "Error",
        description: "Failed to load doctors. Using demo data instead.",
      });
      setUseMockData(true);
    }
  }, [error, toast]);

  if (isLoading && !useMockData) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading doctors...</div>
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
          {useMockData && (
            <div className="mt-4 p-2 bg-yellow-50 border border-yellow-200 rounded-md inline-block">
              <p className="text-yellow-700 text-sm">
                Demo Mode: Displaying example doctors data. API is not available.
              </p>
            </div>
          )}
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
