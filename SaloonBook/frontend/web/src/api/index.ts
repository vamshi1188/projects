import { apiClient } from './client';
import type { Service, Booking, HealthResponse } from '@/types';

// Mock services data for frontend-only application
const mockServices: Service[] = [
  {
    id: 1,
    name: 'Haircut',
    durationMinutes: 30,
    price: 500,
    description: 'Professional haircut services',
  },
  {
    id: 2,
    name: 'Beard',
    durationMinutes: 20,
    price: 250,
    description: 'Beard trimming and styling',
  },
  {
    id: 3,
    name: 'Hair & Beard Color',
    durationMinutes: 60,
    price: 2800,
    description: 'Professional hair and beard coloring',
  },
  {
    id: 4,
    name: 'Massage',
    durationMinutes: 45,
    price: 600,
    description: 'Relaxing massage services',
  },
  {
    id: 5,
    name: 'Face Wash',
    durationMinutes: 25,
    price: 450,
    description: 'Facial cleansing treatments',
  },
];

export const api = {
  health: {
    check: async (): Promise<HealthResponse> => {
      return { status: 'healthy', version: '1.0.0' };
    },
  },

  services: {
    list: async (): Promise<Service[]> => {
      // Return mock services data
      return new Promise((resolve) => {
        setTimeout(() => resolve(mockServices), 100);
      });
    },
  },

  bookings: {
    list: () => apiClient.get<Booking[]>('bookings'),
    create: (booking: Booking) => apiClient.post<Booking>('bookings', booking),
  },
};
