FRONTEND_DIR=frontend/web

.PHONY: install dev build clean docker serve test

install:
	cd $(FRONTEND_DIR) && npm install

dev:
	cd $(FRONTEND_DIR) && npm run dev

build:
	cd $(FRONTEND_DIR) && npm install && npm run build

clean:
	rm -rf $(FRONTEND_DIR)/dist $(FRONTEND_DIR)/node_modules

docker: build
	docker compose up -d

serve: build
	cd $(FRONTEND_DIR) && npx serve dist

test:
	cd $(FRONTEND_DIR) && npm run typecheck
