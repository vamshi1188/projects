FRONTEND_DIR=frontend/web
BACKEND_DIR=backend
PORT?=5000
DB_URL?=postgres://postgres:postgres@localhost:5433/saloonbook_dev

.PHONY: dev backend frontend build clean docker test

dev: backend frontend

backend:
	cd $(BACKEND_DIR) && PORT=$(PORT) DATABASE_URL=$(DB_URL) go run ./cmd

frontend:
	cd $(FRONTEND_DIR) && npm run dev

build:
	cd $(FRONTEND_DIR) && npm install && npm run build
	cd $(BACKEND_DIR) && go build -o bin/saloonbook ./cmd

clean:
	rm -rf $(FRONTEND_DIR)/dist $(BACKEND_DIR)/bin

docker:
	docker compose up --build -d db backend

test:
	cd $(BACKEND_DIR) && go test ./...
	cd $(FRONTEND_DIR) && npm run typecheck
