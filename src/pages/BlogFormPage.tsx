
import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { blogService } from '../services/api';
import Layout from '@/components/Layout';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { useToast } from '@/components/ui/use-toast';
import { useAuth } from '../context/AuthContext';
import { FileText } from 'lucide-react';
import type { Blog } from '@/types/models';

export default function BlogFormPage() {
  const { blogId } = useParams();
  const isEditMode = !!blogId;
  const navigate = useNavigate();
  const { toast } = useToast();
  const { currentUser, hasRole } = useAuth();
  const queryClient = useQueryClient();
  
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [excerpt, setExcerpt] = useState('');
  const [coverImage, setCoverImage] = useState('');
  const [tags, setTags] = useState('');
  
  // Check permission
  useEffect(() => {
    if (!hasRole(['admin', 'superadmin', 'doctor'])) {
      toast({
        variant: "destructive",
        title: "Access Denied",
        description: "You don't have permission to create or edit blog posts",
      });
      navigate('/blogs');
    }
  }, [hasRole, navigate, toast]);
  
  // Fetch blog data if in edit mode
  const { data: blog, isLoading } = useQuery({
    queryKey: ['blog', blogId],
    queryFn: () => blogService.getBlogById(Number(blogId)),
    enabled: isEditMode,
    onSuccess: (data) => {
      // Check if current user is allowed to edit this blog
      const canEdit = hasRole(['admin', 'superadmin']) || 
        (hasRole(['doctor']) && currentUser?.id === data.authorId);
        
      if (!canEdit) {
        toast({
          variant: "destructive",
          title: "Access Denied",
          description: "You don't have permission to edit this blog post",
        });
        navigate(`/blogs/${blogId}`);
        return;
      }
      
      setTitle(data.title);
      setContent(data.content);
      setExcerpt(data.excerpt || '');
      setCoverImage(data.coverImage || '');
      setTags(data.tags ? data.tags.join(', ') : '');
    },
  });
  
  // Create blog mutation
  const createMutation = useMutation({
    mutationFn: (newBlog: Partial<Blog>) => blogService.createBlog(newBlog),
    onSuccess: (data) => {
      toast({
        title: "Blog created",
        description: "Your blog post has been published successfully",
      });
      queryClient.invalidateQueries({ queryKey: ['blogs'] });
      navigate(`/blogs/${data.id}`);
    },
    onError: () => {
      toast({
        variant: "destructive",
        title: "Error",
        description: "Failed to create blog post. Please try again.",
      });
    },
  });
  
  // Update blog mutation
  const updateMutation = useMutation({
    mutationFn: (updatedBlog: Partial<Blog>) => 
      blogService.updateBlog(Number(blogId), updatedBlog),
    onSuccess: (data) => {
      toast({
        title: "Blog updated",
        description: "Your blog post has been updated successfully",
      });
      queryClient.invalidateQueries({ queryKey: ['blogs'] });
      queryClient.invalidateQueries({ queryKey: ['blog', blogId] });
      navigate(`/blogs/${data.id}`);
    },
    onError: () => {
      toast({
        variant: "destructive",
        title: "Error",
        description: "Failed to update blog post. Please try again.",
      });
    },
  });
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!title.trim() || !content.trim()) {
      toast({
        variant: "destructive",
        title: "Missing information",
        description: "Title and content are required",
      });
      return;
    }
    
    const blogData = {
      title: title.trim(),
      content: content.trim(),
      excerpt: excerpt.trim() || null,
      coverImage: coverImage.trim() || null,
      tags: tags.trim() ? tags.split(',').map(tag => tag.trim()) : [],
      authorId: currentUser?.id,
      authorName: currentUser?.name,
    };
    
    if (isEditMode) {
      updateMutation.mutate(blogData);
    } else {
      createMutation.mutate(blogData);
    }
  };
  
  if (isLoading) {
    return (
      <Layout>
        <div className="flex justify-center items-center h-64">
          <div className="animate-pulse text-xl">Loading blog post...</div>
        </div>
      </Layout>
    );
  }
  
  return (
    <Layout>
      <div className="container max-w-4xl mx-auto py-8 px-4">
        <Card>
          <CardHeader>
            <CardTitle>{isEditMode ? 'Edit Blog Post' : 'Create New Blog Post'}</CardTitle>
            <CardDescription>
              {isEditMode 
                ? 'Update your existing blog post' 
                : 'Share your knowledge with the CareHub community'}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-2">
                <Label htmlFor="title">Title *</Label>
                <Input
                  id="title"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  placeholder="Enter blog title"
                  required
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="excerpt">Short Excerpt (Optional)</Label>
                <Textarea
                  id="excerpt"
                  value={excerpt}
                  onChange={(e) => setExcerpt(e.target.value)}
                  placeholder="Brief description of your post (shown in previews)"
                  rows={3}
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="content">Content *</Label>
                <Textarea
                  id="content"
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="Write your blog content here"
                  rows={12}
                  required
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="coverImage">Cover Image URL (Optional)</Label>
                <Input
                  id="coverImage"
                  value={coverImage}
                  onChange={(e) => setCoverImage(e.target.value)}
                  placeholder="Enter image URL for blog cover"
                />
                {coverImage && (
                  <div className="aspect-video w-full overflow-hidden rounded-md mt-2">
                    <img 
                      src={coverImage} 
                      alt="Cover preview" 
                      className="h-full w-full object-cover"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.src = "https://placehold.co/600x400?text=Invalid+Image+URL";
                      }}
                    />
                  </div>
                )}
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="tags">Tags (Optional, comma-separated)</Label>
                <Input
                  id="tags"
                  value={tags}
                  onChange={(e) => setTags(e.target.value)}
                  placeholder="e.g. cardiology, wellness, research"
                />
              </div>
              
              <div className="pt-4 flex justify-between">
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => navigate(isEditMode ? `/blogs/${blogId}` : '/blogs')}
                >
                  Cancel
                </Button>
                
                <Button 
                  type="submit"
                  disabled={createMutation.isPending || updateMutation.isPending}
                >
                  {(createMutation.isPending || updateMutation.isPending) ? (
                    <span>Saving...</span>
                  ) : (
                    <>
                      <FileText className="mr-2 h-4 w-4" />
                      {isEditMode ? 'Update Post' : 'Publish Post'}
                    </>
                  )}
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </Layout>
  );
}
