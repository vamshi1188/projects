# SaloonBook Frontend

React + TypeScript + Vite frontend for SaloonBook.

## Structure

```
frontend/web/
├── src/
│   ├── api/             # API client and services
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

## Build

```bash
npm run build
```

## Environment Variables

Create `.env.development`:

```
VITE_API_URL=http://localhost:5000/api
```
