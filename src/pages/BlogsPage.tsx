
import React from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { blogService } from '../services/api';
import Layout from '@/components/Layout';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { useToast } from '@/components/ui/use-toast';
import { Pencil, Calendar, User } from 'lucide-react';
import { useAuth } from '../context/AuthContext';
import type { Blog } from '@/types/models';

export default function BlogsPage() {
  const { toast } = useToast();
  const { currentUser, hasRole } = useAuth();
  
  const { data: blogs, isLoading, error } = useQuery({
    queryKey: ['blogs'],
    queryFn: blogService.getAllBlogs,
  });

  if (isLoading) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading blogs...</div>
        </div>
      </Layout>
    );
  }

  if (error) {
    toast({
      variant: "destructive",
      title: "Error",
      description: "Failed to load blog posts",
    });
    
    return (
      <Layout>
        <div className="text-center py-10">
          <h2 className="text-xl font-bold">Error loading blog posts</h2>
          <p className="text-muted-foreground">Please try again later</p>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="container max-w-6xl mx-auto py-8">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold">CareHub Blog</h1>
          
          {hasRole(['admin', 'superadmin', 'doctor']) && (
            <Link to="/blogs/new">
              <Button className="gap-2">
                <Pencil className="h-4 w-4" />
                Write New Post
              </Button>
            </Link>
          )}
        </div>
        
        {blogs && blogs.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {blogs.map((blog: Blog) => (
              <Card key={blog.id} className="flex flex-col h-full overflow-hidden">
                {blog.coverImage && (
                  <div className="aspect-video w-full overflow-hidden">
                    <img 
                      src={blog.coverImage} 
                      alt={blog.title} 
                      className="h-full w-full object-cover transition-transform hover:scale-105"
                    />
                  </div>
                )}
                <CardHeader className="pb-4">
                  <CardTitle>
                    <Link to={`/blogs/${blog.id}`} className="hover:text-primary">
                      {blog.title}
                    </Link>
                  </CardTitle>
                  <CardDescription className="flex flex-col gap-1 text-xs">
                    <div className="flex items-center gap-1">
                      <User className="h-3 w-3" />
                      <Link 
                        to={`/doctors/${blog.authorId}`}
                        className="hover:text-primary"
                      >
                        {blog.authorName}
                      </Link>
                    </div>
                    <div className="flex items-center gap-1">
                      <Calendar className="h-3 w-3" />
                      {new Date(blog.publishedAt).toLocaleDateString()}
                    </div>
                  </CardDescription>
                </CardHeader>
                <CardContent className="flex-grow">
                  <p className="line-clamp-3 text-sm text-muted-foreground">
                    {blog.excerpt || blog.content.substring(0, 150)}...
                  </p>
                </CardContent>
                <CardFooter>
                  <Link 
                    to={`/blogs/${blog.id}`}
                    className="text-primary text-sm font-medium hover:underline"
                  >
                    Read more
                  </Link>
                </CardFooter>
              </Card>
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            <h2 className="text-xl font-semibold mb-2">No blog posts yet</h2>
            <p className="text-muted-foreground">
              {hasRole(['admin', 'superadmin', 'doctor']) 
                ? "Why not write the first one?" 
                : "Check back soon for new articles."}
            </p>
            
            {hasRole(['admin', 'superadmin', 'doctor']) && (
              <Link to="/blogs/new">
                <Button className="mt-4">
                  <Pencil className="mr-2 h-4 w-4" />
                  Write New Post
                </Button>
              </Link>
            )}
          </div>
        )}
      </div>
    </Layout>
  );
}
