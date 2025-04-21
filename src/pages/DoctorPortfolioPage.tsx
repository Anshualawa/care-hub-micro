
import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { doctorService, blogService } from '../services/api';
import Layout from '@/components/Layout';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import { useToast } from '@/components/ui/use-toast';
import { Pencil, GraduationCap, Briefcase, FileText } from 'lucide-react';

export default function DoctorPortfolioPage() {
  const { doctorId } = useParams();
  const { toast } = useToast();
  
  const {
    data: doctor,
    isLoading: isLoadingDoctor,
    error: doctorError
  } = useQuery({
    queryKey: ['doctor', doctorId],
    queryFn: () => doctorService.getDoctorById(Number(doctorId)),
  });

  const {
    data: blogs,
    isLoading: isLoadingBlogs,
    error: blogsError
  } = useQuery({
    queryKey: ['doctorBlogs', doctorId],
    queryFn: () => blogService.getBlogsByAuthor(Number(doctorId)),
  });

  if (isLoadingDoctor || isLoadingBlogs) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading doctor profile...</div>
        </div>
      </Layout>
    );
  }

  if (doctorError) {
    toast({
      variant: "destructive",
      title: "Error",
      description: "Failed to load doctor profile",
    });
    
    return (
      <Layout>
        <div className="text-center py-10">
          <h2 className="text-xl font-bold">Error loading doctor profile</h2>
          <p className="text-muted-foreground">Please try again later</p>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="container mx-auto py-8">
        <div className="flex flex-col md:flex-row gap-8">
          {/* Left column - Doctor profile */}
          <div className="w-full md:w-1/3">
            <Card>
              <CardContent className="pt-6 flex flex-col items-center">
                <Avatar className="w-32 h-32 mb-4 border-2 border-primary">
                  <AvatarImage src={doctor?.profileImage} />
                  <AvatarFallback className="bg-primary/10 text-primary text-2xl">
                    {doctor?.name?.split(' ').map((n: string) => n[0]).join('')}
                  </AvatarFallback>
                </Avatar>
                
                <h2 className="text-2xl font-bold">{doctor?.name}</h2>
                <p className="text-muted-foreground">{doctor?.specialization}</p>
                <Badge className="mt-2">{doctor?.department}</Badge>
                
                <Separator className="my-4" />
                
                <div className="w-full">
                  <div className="flex gap-2 items-center mb-2">
                    <FileText className="text-muted-foreground h-4 w-4" />
                    <span className="text-muted-foreground">Email:</span>
                    <span>{doctor?.email}</span>
                  </div>
                  <div className="flex gap-2 items-center">
                    <FileText className="text-muted-foreground h-4 w-4" />
                    <span className="text-muted-foreground">Phone:</span>
                    <span>{doctor?.phone}</span>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
          
          {/* Right column - Tabs with details */}
          <div className="w-full md:w-2/3">
            <Tabs defaultValue="about">
              <TabsList className="mb-4">
                <TabsTrigger value="about">About</TabsTrigger>
                <TabsTrigger value="education">Education</TabsTrigger>
                <TabsTrigger value="experience">Experience</TabsTrigger>
                <TabsTrigger value="blogs">Blogs</TabsTrigger>
              </TabsList>
              
              <TabsContent value="about">
                <Card>
                  <CardHeader>
                    <CardTitle>About</CardTitle>
                    <CardDescription>Doctor's biography and specialization</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p className="whitespace-pre-wrap">{doctor?.bio || "No biography available."}</p>
                  </CardContent>
                </Card>
              </TabsContent>
              
              <TabsContent value="education">
                <Card>
                  <CardHeader>
                    <CardTitle>Education</CardTitle>
                    <CardDescription>Academic background and qualifications</CardDescription>
                  </CardHeader>
                  <CardContent>
                    {doctor?.education && doctor.education.length > 0 ? (
                      <ul className="space-y-4">
                        {doctor.education.map((edu: string, index: number) => (
                          <li key={index} className="flex gap-2">
                            <GraduationCap className="h-5 w-5 text-primary flex-shrink-0" />
                            <span>{edu}</span>
                          </li>
                        ))}
                      </ul>
                    ) : (
                      <p>No education information available.</p>
                    )}
                  </CardContent>
                </Card>
              </TabsContent>
              
              <TabsContent value="experience">
                <Card>
                  <CardHeader>
                    <CardTitle>Experience</CardTitle>
                    <CardDescription>Professional experience and career history</CardDescription>
                  </CardHeader>
                  <CardContent>
                    {doctor?.experience && doctor.experience.length > 0 ? (
                      <ul className="space-y-4">
                        {doctor.experience.map((exp: string, index: number) => (
                          <li key={index} className="flex gap-2">
                            <Briefcase className="h-5 w-5 text-primary flex-shrink-0" />
                            <span>{exp}</span>
                          </li>
                        ))}
                      </ul>
                    ) : (
                      <p>No experience information available.</p>
                    )}
                  </CardContent>
                </Card>
              </TabsContent>
              
              <TabsContent value="blogs">
                <Card>
                  <CardHeader>
                    <CardTitle>Blog Posts</CardTitle>
                    <CardDescription>Articles written by this doctor</CardDescription>
                  </CardHeader>
                  <CardContent>
                    {blogs && blogs.length > 0 ? (
                      <div className="space-y-4">
                        {blogs.map((blog: any) => (
                          <Card key={blog.id} className="overflow-hidden">
                            {blog.coverImage && (
                              <img 
                                src={blog.coverImage} 
                                alt={blog.title} 
                                className="h-32 w-full object-cover"
                              />
                            )}
                            <CardHeader className="pb-2">
                              <CardTitle className="text-lg">
                                <Link to={`/blogs/${blog.id}`} className="hover:text-primary">
                                  {blog.title}
                                </Link>
                              </CardTitle>
                              <CardDescription className="flex items-center gap-1">
                                <Pencil className="h-3 w-3" />
                                Published on {new Date(blog.publishedAt).toLocaleDateString()}
                              </CardDescription>
                            </CardHeader>
                            <CardContent>
                              <p className="line-clamp-3">{blog.excerpt || blog.content.substring(0, 150)}...</p>
                              <Link 
                                to={`/blogs/${blog.id}`}
                                className="text-primary text-sm font-medium mt-2 inline-block hover:underline"
                              >
                                Read more
                              </Link>
                            </CardContent>
                          </Card>
                        ))}
                      </div>
                    ) : (
                      <p>No blog posts available from this doctor.</p>
                    )}
                  </CardContent>
                </Card>
              </TabsContent>
            </Tabs>
          </div>
        </div>
      </div>
    </Layout>
  );
}
