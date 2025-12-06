## Running SaloonBook (Frontend-Only App)

### No Backend or Database Required! 🎉

SaloonBook is now a **100% frontend application**. All data is stored in your browser's LocalStorage.

### Quick Start

```bash
cd frontend/web
npm install
npm run dev
```

Then open http://localhost:5173

### Using Makefile

```bash
# Install dependencies
make install

# Run development server
make dev

# Build for production
make build
```

### Production Deployment

Build the static files:

```bash
cd frontend/web
npm run build
```

Serve with any static file server:

```bash
# Option 1: Using serve
npx serve dist

# Option 2: Using Python
python -m http.server 8080 -d dist

# Option 3: Using nginx with Docker
make docker
```

### Verify It's Working

1. Open http://localhost:5173 (dev) or http://localhost:8080 (prod)
2. Click "Get Started" on the welcome screen
3. Select services and styles
4. Create a booking
5. Check your browser's LocalStorage to see the saved data:
   - Open DevTools (F12)
   - Go to Application > Local Storage
   - Look for `bookings` key

### Features

✅ **No Backend Required** - Runs 100% in browser  
✅ **LocalStorage** - Data persists across sessions  
✅ **Service Selection** - Multiple service types and styles  
✅ **Booking Management** - Create and view bookings  
✅ **QR Codes** - Generate booking confirmations  
✅ **Responsive Design** - Mobile-first UI
