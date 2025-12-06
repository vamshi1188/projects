# Quick Start Guide

## Frontend-Only Application

SaloonBook is now a **100% frontend application** with no backend or database required. All data is stored in your browser's LocalStorage.

## Start the Application

### Install Dependencies
```bash
cd frontend/web
npm install
```

### Development Mode
```bash
cd frontend/web
npm run dev
```

Then open http://localhost:5173 in your browser.

## Alternative: Using Makefile

```bash
# Install dependencies
make install

# Run development server
make dev

# Build for production
make build

# Serve production build
make serve
```

## Production Build

Build the optimized production bundle:

```bash
cd frontend/web
npm run build
```

The output will be in `frontend/web/dist/`.

Serve it with any static file server:

```bash
# Using npx serve
npx serve dist

# Using Python
python -m http.server 8080 -d dist

# Using nginx with Docker
make docker
```

## Features

- ✅ No backend required
- ✅ No database setup needed
- ✅ All data stored in browser LocalStorage
- ✅ Service selection and booking management
- ✅ QR code generation
- ✅ Responsive mobile-first design

## Test the Application

Once running:
- Frontend: http://localhost:5173
- Backend Health: http://localhost:5000/api/health
- Backend Services: http://localhost:5000/api/services

## If You Want to Use Port 5432

Stop your local PostgreSQL first:
```bash
# macOS
brew services stop postgresql

# Linux (systemd)
sudo systemctl stop postgresql

# Then revert docker-compose.yml to use 5432
```

## Current Setup

✅ Docker PostgreSQL: port **5433** → 5432 (internal)  
✅ Backend API: port **5000**  
✅ Frontend Dev: port **5173**

Everything is configured and ready to run!
