#!/bin/bash
# Development run script for backend

cd "$(dirname "$0")"

# Check if .env exists
if [ ! -f .env ]; then
    echo "⚠️  .env file not found, creating from example..."
    cat > .env << EOF
DATABASE_URL=postgres://postgres:postgres@localhost:5433/saloonbook_dev
PORT=5000
NODE_ENV=development
SESSION_SECRET=dev-secret
SERVE_FRONTEND=0
EOF
fi

# Load environment variables
export $(cat .env | grep -v '^#' | xargs)

echo "🚀 Starting backend..."
echo "   Database: $DATABASE_URL"
echo "   Port: $PORT"
echo ""

go run ./cmd
