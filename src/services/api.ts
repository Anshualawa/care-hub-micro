
import axios from 'axios';

const API_URL = 'http://localhost:8090';

// Create axios instance
const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add authentication interceptor
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth services
export const authService = {
  login: async (username: string, password: string) => {
    const response = await api.post('/auth/login', { username, password });
    if (response.data.token) {
      localStorage.setItem('token', response.data.token);
      localStorage.setItem('user', JSON.stringify(response.data.user));
    }
    return response.data;
  },
  logout: () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  },
  getCurrentUser: () => {
    const user = localStorage.getItem('user');
    return user ? JSON.parse(user) : null;
  }
};

// Patient services
export const patientService = {
  getAllPatients: async () => {
    const response = await api.get('/api/patients');
    return response.data;
  },
  getPatient: async (id: number) => {
    const response = await api.get(`/api/patients/${id}`);
    return response.data;
  },
  createPatient: async (patient: any) => {
    const response = await api.post('/api/patients', patient);
    return response.data;
  },
  updatePatient: async (id: number, patient: any) => {
    const response = await api.put(`/api/patients/${id}`, patient);
    return response.data;
  },
  deletePatient: async (id: number) => {
    const response = await api.delete(`/api/patients/${id}`);
    return response.data;
  }
};

// Appointment services
export const appointmentService = {
  getAllAppointments: async () => {
    const response = await api.get('/api/appointments');
    return response.data;
  },
  getPatientAppointments: async (patientId: number) => {
    const response = await api.get(`/api/appointments?patientId=${patientId}`);
    return response.data;
  },
  getAppointment: async (id: number) => {
    const response = await api.get(`/api/appointments/${id}`);
    return response.data;
  },
  createAppointment: async (appointment: any) => {
    const response = await api.post('/api/appointments', appointment);
    return response.data;
  },
  updateAppointment: async (id: number, appointment: any) => {
    const response = await api.put(`/api/appointments/${id}`, appointment);
    return response.data;
  },
  deleteAppointment: async (id: number) => {
    const response = await api.delete(`/api/appointments/${id}`);
    return response.data;
  }
};

// Health metric services
export const healthMetricService = {
  getPatientMetrics: async (patientId: number) => {
    const response = await api.get(`/api/patients/${patientId}/metrics`);
    return response.data;
  },
  recordMetric: async (patientId: number, metric: any) => {
    const response = await api.post(`/api/patients/${patientId}/metrics`, metric);
    return response.data;
  }
};

export default api;
