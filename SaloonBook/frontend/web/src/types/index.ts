export interface Service {
  id: number;
  name: string;
  durationMinutes: number;
  price: number;
  description?: string;
}

export interface Booking {
  id?: number;
  serviceId: number;
  customer: string;
  phone: string;
  status?: string;
  createdAt?: string;
}

export interface HealthResponse {
  status: string;
  version?: string;
}

export type ServiceType = 'haircut' | 'beard' | 'color' | 'massage' | 'facewash';

export interface Style {
  id: string;
  name: string;
  price: number;
  image: string;
}

export interface SelectedService {
  serviceType: ServiceType;
  serviceName: string;
  styles: Style[];
}
