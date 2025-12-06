# SaloonBook Frontend

React + TypeScript + Vite frontend-only salon booking application.

## Structure

```
frontend/web/
├── src/
│   ├── api/             # LocalStorage-based data layer
│   ├── components/      # React components
│   │   ├── ui/          # Reusable UI components
│   │   └── examples/    # Example components
│   ├── hooks/           # Custom React hooks
│   ├── lib/             # Utilities and helpers
│   ├── pages/           # Page components
│   ├── types/           # TypeScript types
│   ├── constants/       # Constants and config
│   └── utils/           # Utility functions
├── public/              # Static assets
└── docs/                # Documentation
```

## Development

```bash
cd frontend/web
npm install
npm run dev
```

Open http://localhost:5173 in your browser.

## Build

```bash
npm run build
```

The optimized production build will be in `dist/`.

## Features

- **100% Frontend**: No backend or database required
- **LocalStorage**: All data persisted in browser
- **Service Selection**: Haircut, Beard, Color, Massage, Face Wash
- **Phone Verification**: OTP-style verification flow
- **Booking Management**: View and manage bookings
- **QR Codes**: Generate booking confirmation QR codes
- **Responsive Design**: Mobile-first UI with Tailwind CSS

## No Backend Required

This app runs entirely in your browser using LocalStorage for data persistence. No server, database, or API required!
