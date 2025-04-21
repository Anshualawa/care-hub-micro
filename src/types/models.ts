
export interface Patient {
  id: number;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  email?: string;
  phone?: string;
  address?: string;
  createdAt: string;
}

export interface Appointment {
  id: number;
  patientId: number;
  dateTime: string;
  description: string;
  status: string;
  doctor: string;
}

export interface Doctor {
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

export interface Blog {
  id: number;
  title: string;
  content: string;
  excerpt?: string;
  coverImage?: string;
  authorId: number;
  authorName: string;
  publishedAt: string;
  updatedAt?: string;
  tags?: string[];
}
