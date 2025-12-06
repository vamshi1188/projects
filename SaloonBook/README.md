# SaloonBook

> React (Vite) frontend-only salon booking application with Android APK support.

---

## Quick overview

- Frontend: `frontend/web/` — Vite + React + TypeScript UI
- All data is stored locally in the browser (LocalStorage)
- No backend or database required
- **NEW:** Can be built as an Android APK using Capacitor

This README shows the most common commands to develop, build and run the app locally.

## 📱 Build Android APK

🎉 **NEW!** Your app is ready to build as an Android APK!

### Quick Start (3 Steps):
```bash
# 1. Check environment
./check-android-env.sh

# 2. Run setup
./setup-android.sh

# 3. Build APK
cd frontend/web && npx cap open android
```

### 📚 Complete Documentation:
- **[BUILD_ANDROID_README.md](./BUILD_ANDROID_README.md)** - 📖 **START HERE** - Complete overview
- **[QUICKSTART_ANDROID.md](./QUICKSTART_ANDROID.md)** - ⚡ Quick commands reference
- **[ANDROID_BUILD_GUIDE.md](./ANDROID_BUILD_GUIDE.md)** - 📘 Detailed step-by-step guide
- **[ANDROID_BUILD_FLOW.md](./ANDROID_BUILD_FLOW.md)** - 🎯 Visual diagrams and flows
- **[ANDROID_CHECKLIST.md](./ANDROID_CHECKLIST.md)** - ✅ Verification checklist
- **[ANDROID_TROUBLESHOOTING.md](./ANDROID_TROUBLESHOOTING.md)** - 🔧 Problem solving guide

Your APK will be at: `frontend/web/android/app/build/outputs/apk/debug/app-debug.apk`

## Prerequisites

- Node.js 18+ (recommend latest LTS)
- npm (or yarn / pnpm)

Optional (for future SMS / payments integration): Twilio credentials, Stripe keys.

## Install dependencies

Use your preferred package manager. Examples below use npm.

Install (one-time):

```bash
cd frontend/web
npm install
```

If you prefer pnpm or yarn:

```bash
# pnpm
cd frontend/web
pnpm install

# yarn
cd frontend/web
yarn install
```

## Development

Run the development server:

```bash
cd frontend/web
npm run dev
```

Open http://localhost:5173

The app runs entirely in the browser with no backend required. All bookings and data are stored in your browser's LocalStorage.

## Production build

Build the optimized React app:

```bash
cd frontend/web
npm run build
```

The build output will be in `frontend/web/dist/`. You can serve this with any static file server:

```bash
# Using npx serve
npx serve dist

# Using Python
python -m http.server 8080 -d dist

# Using Node.js http-server
npx http-server dist
```

## Features

- Service selection (Haircut, Beard, Hair & Beard Color, Massage, Face Wash)
- Multiple style options for each service
- Phone verification flow
- Booking management
- Order summary and confirmation
- QR code generation for bookings
- Responsive mobile-first design

## Project Structure

```
frontend/web/
├── src/
│   ├── components/     # React components
│   ├── api/           # API client (now using LocalStorage)
│   ├── constants/     # App constants
│   ├── types/         # TypeScript types
│   ├── hooks/         # Custom React hooks
│   └── lib/           # Utility functions
├── public/            # Static assets
└── attached_assets/   # Images and generated assets
```

---

© SaloonBook
