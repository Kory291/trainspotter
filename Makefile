.PHONY: build

FRONTEND_DIR = ./frontend
BACKEND_DIR = ./backend
VERSION = v1

build:
	docker build -f $(FRONTEND_DIR)/Dockerfile -t trainspotter-frontend:$(VERSION) $(FRONTEND_DIR)
	docker build -f $(BACKEND_DIR)/Dockerfile -t trainspotter-backend:$(VERSION) $(BACKEND_DIR)
