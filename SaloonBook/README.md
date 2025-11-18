# SaloonBook

> Go backend + React (Vite) frontend.

---

## Quick overview

- Frontend: `frontend/web/` — Vite + React + TypeScript UI
- Backend: `backend/` — Go (chi router, pgx Postgres)
- DB: PostgreSQL via `DATABASE_URL`

This README shows the most common commands to develop, build and run the app locally.

## Prerequisites

- Node.js 18+ (recommend latest LTS)
- npm (or yarn / pnpm)
- A Postgres-compatible database (Neon, Supabase, local Postgres, etc.)

Optional (for SMS / payments): Twilio credentials, Stripe keys.

## Environment variables

Create a `.env` file at the project root (or set environment variables in your shell). Minimal variables the app expects:

```
DATABASE_URL=postgres://user:pass@host:5432/dbname
SESSION_SECRET=your-session-secret
PORT=5000
NODE_ENV=development
# Optional (if you add payments or SMS later):
# STRIPE_SECRET_KEY=sk_live_...
# STRIPE_PUBLISHABLE_KEY=pk_live_...
# TWILIO_ACCOUNT_SID=...
# TWILIO_AUTH_TOKEN=...
```

Keep secrets out of source control. Use a vault or CI secrets for production.

## Install dependencies

Use your preferred package manager. Examples below use npm.

Install (one-time):

```bash
npm install
```

If you prefer pnpm or yarn:

```bash
# pnpm
pnpm install

# yarn
yarn install
```

## Setup

```bash
# Frontend
cd frontend/web
npm install

# Backend
cd backend
go mod tidy
```

## Development

Terminal 1:
```bash
cd backend
PORT=5000 DATABASE_URL=postgres://postgres:postgres@localhost:5432/saloonbook_dev go run ./cmd/saloonbook
```

Terminal 2:
```bash
cd frontend/web
npm run dev
```

Open http://localhost:5173

API proxied at `/api/*`.

Health check:

```bash
curl http://localhost:5000/api/health
```

Example services:

```bash
curl http://localhost:5000/api/services
```

Create booking:

```bash
curl -X POST http://localhost:5000/api/bookings -H 'Content-Type: application/json' \
 -d '{"serviceId":1,"customer":"Alice","phone":"123456"}'
```

List bookings:

```bash
curl http://localhost:5000/api/bookings
```

## Production build (frontend)

Build React:

```bash
cd frontend/web
npm run build
```

(Optional) Serve `dist/` via nginx or add static file serving in Go.

## Cleaning

Old Node/Express artifacts removed (`dist/`, drizzle config). Use a Go migration tool before real DB writes.

## Notes

Original Node/Express backend replaced by Go; Drizzle migrations not yet ported—add a Go migration tool (e.g. goose, sqlc) next.

---

© SaloonBook
