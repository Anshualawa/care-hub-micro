
import React from 'react';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { blogService } from '../services/api';
import Layout from '@/components/Layout';
import { Button } from '@/components/ui/button';
import { useToast } from '@/components/ui/use-toast';
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar';
import { Badge } from '@/components/ui/badge';
import { Calendar, Clock, User, Edit, Pencil } from 'lucide-react';
import { useAuth } from '../context/AuthContext';

export default function BlogDetailPage() {
  const { blogId } = useParams();
  const navigate = useNavigate();
  const { toast } = useToast();
  const { currentUser, hasRole } = useAuth();
  
  const { data: blog, isLoading, error } = useQuery({
    queryKey: ['blog', blogId],
    queryFn: () => blogService.getBlogById(Number(blogId)),
  });

  if (isLoading) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading blog post...</div>
        </div>
      </Layout>
    );
  }

  if (error || !blog) {
    toast({
      variant: "destructive",
      title: "Error",
      description: "Failed to load blog post",
    });
    
    return (
      <Layout>
        <div className="text-center py-10">
          <h2 className="text-xl font-bold">Error loading blog post</h2>
          <p className="text-muted-foreground">Please try again later</p>
          <Button 
            variant="outline" 
            onClick={() => navigate('/blogs')}
            className="mt-4"
          >
            Back to Blogs
          </Button>
        </div>
      </Layout>
    );
  }

  const canEdit = hasRole(['admin', 'superadmin']) || 
    (hasRole(['doctor']) && currentUser?.id === blog.authorId);

  return (
    <Layout>
      <div className="container max-w-4xl mx-auto py-8 px-4">
        <div className="mb-8">
          <Link 
            to="/blogs"
            className="text-sm text-muted-foreground hover:text-primary mb-4 inline-block"
          >
            ← Back to all blogs
          </Link>
          
          {blog.coverImage && (
            <div className="aspect-video w-full overflow-hidden rounded-lg mb-6">
              <img 
                src={blog.coverImage} 
                alt={blog.title} 
                className="h-full w-full object-cover"
              />
            </div>
          )}
          
          <h1 className="text-3xl md:text-4xl font-bold mb-4">{blog.title}</h1>
          
          <div className="flex flex-wrap items-center gap-4 mb-6">
            <div className="flex items-center gap-2">
              <Avatar className="h-8 w-8">
                <AvatarFallback>
                  {blog.authorName?.split(' ').map(n => n[0]).join('')}
                </AvatarFallback>
              </Avatar>
              <div>
                <Link 
                  to={`/doctors/${blog.authorId}`}
                  className="text-sm font-medium hover:text-primary"
                >
                  {blog.authorName}
                </Link>
              </div>
            </div>
            
            <div className="flex items-center gap-1 text-sm text-muted-foreground">
              <Calendar className="h-4 w-4" />
              <span>{new Date(blog.publishedAt).toLocaleDateString()}</span>
            </div>
            
            {blog.updatedAt && blog.updatedAt !== blog.publishedAt && (
              <div className="flex items-center gap-1 text-sm text-muted-foreground">
                <Clock className="h-4 w-4" />
                <span>Updated {new Date(blog.updatedAt).toLocaleDateString()}</span>
              </div>
            )}
            
            {canEdit && (
              <Link to={`/blogs/edit/${blog.id}`} className="ml-auto">
                <Button variant="outline" size="sm">
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Post
                </Button>
              </Link>
            )}
          </div>
          
          {blog.tags && blog.tags.length > 0 && (
            <div className="flex flex-wrap gap-2 mb-6">
              {blog.tags.map((tag, index) => (
                <Badge key={index} variant="secondary">{tag}</Badge>
              ))}
            </div>
          )}
        </div>
        
        <div className="prose prose-sm sm:prose lg:prose-lg max-w-none">
          <div dangerouslySetInnerHTML={{ __html: blog.content }} />
        </div>
        
        <div className="mt-12 pt-8 border-t">
          <h3 className="text-lg font-semibold mb-4">About the Author</h3>
          <Link to={`/doctors/${blog.authorId}`} className="flex items-start gap-4">
            <Avatar className="h-12 w-12">
              <AvatarFallback className="bg-primary/10 text-primary">
                {blog.authorName?.split(' ').map(n => n[0]).join('')}
              </AvatarFallback>
            </Avatar>
            <div>
              <p className="font-medium">{blog.authorName}</p>
              <p className="text-sm text-muted-foreground">
                View full profile and other articles
              </p>
            </div>
          </Link>
        </div>
      </div>
    </Layout>
  );
}
