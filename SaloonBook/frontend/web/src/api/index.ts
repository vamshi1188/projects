import { apiClient } from './client';
import type { Service, Booking, HealthResponse } from '@/types';

export const api = {
  health: {
    check: () => apiClient.get<HealthResponse>('/health'),
  },

  services: {
    list: () => apiClient.get<Service[]>('/services'),
  },

  bookings: {
    list: () => apiClient.get<Booking[]>('/bookings'),
    create: (booking: Booking) => apiClient.post<Booking>('/bookings', booking),
  },
};
