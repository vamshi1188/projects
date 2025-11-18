export const APP_NAME = 'RNL - Salon Booking';
export const APP_VERSION = '1.0.0';

export const SERVICE_NAMES: Record<string, string> = {
  haircut: 'Haircut',
  beard: 'Beard',
  color: 'Hair & Beard Color',
  massage: 'Massage',
  facewash: 'Face Wash',
};

export const SERVICE_QUEUE = ['haircut', 'beard', 'color', 'massage', 'facewash'] as const;
