// LocalStorage-based API client for frontend-only application
class ApiClient {
  private getFromStorage<T>(key: string, defaultValue: T): T {
    try {
      const item = localStorage.getItem(key);
      return item ? JSON.parse(item) : defaultValue;
    } catch (error) {
      console.error(`Error reading from localStorage: ${error}`);
      return defaultValue;
    }
  }

  private setInStorage<T>(key: string, value: T): void {
    try {
      localStorage.setItem(key, JSON.stringify(value));
    } catch (error) {
      console.error(`Error writing to localStorage: ${error}`);
    }
  }

  async get<T>(endpoint: string): Promise<T> {
    // Simulate async behavior
    return new Promise((resolve) => {
      setTimeout(() => {
        const key = endpoint.replace(/^\//, '');
        const data = this.getFromStorage<T>(key, [] as T);
        resolve(data);
      }, 100);
    });
  }

  async post<T>(endpoint: string, data: unknown): Promise<T> {
    // Simulate async behavior
    return new Promise((resolve) => {
      setTimeout(() => {
        const key = endpoint.replace(/^\//, '');
        const existing = this.getFromStorage<any[]>(key, []);

        // Add ID and timestamp if it's a new item
        const newItem = {
          ...data,
          id: existing.length > 0 ? Math.max(...existing.map((item: any) => item.id || 0)) + 1 : 1,
          createdAt: new Date().toISOString(),
        };

        const updated = [...existing, newItem];
        this.setInStorage(key, updated);
        resolve(newItem as T);
      }, 100);
    });
  }
}

export const apiClient = new ApiClient();
