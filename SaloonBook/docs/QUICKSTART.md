# Quick Start Guide

## ✅ Database is Running on Port 5433

Your local PostgreSQL was using port 5432, so Docker PostgreSQL now uses port **5433**.

## Start the Application

### Backend
```bash
cd backend
DATABASE_URL=postgres://postgres:postgres@localhost:5433/saloonbook_dev \
PORT=5000 \
go run ./cmd
```

### Frontend (separate terminal)
```bash
cd frontend/web
npm install
npm run dev
```

## Alternative: Using Makefile

The Makefile has been updated with the correct port:

```bash
# Terminal 1: Backend
make backend

# Terminal 2: Frontend  
make frontend
```

## Connection Strings

**Local Development:**
```
DATABASE_URL=postgres://postgres:postgres@localhost:5433/saloonbook_dev
```

**Docker Backend (inside compose network):**
```
DATABASE_URL=postgres://postgres:postgres@db:5432/saloonbook_dev
```
(Inside Docker network, it uses internal port 5432)

## Test Endpoints

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
