SHELL := /bin/bash

.PHONY: build up down sh logs

# Build the dev image
build:
	docker compose build

# Start the container in background
up:
	docker compose up -d

# Stop and remove the container
down:
	docker compose down

# Open a shell inside the container (as user 'dev')
sh:
	docker compose exec dev bash

# Follow logs (useful if you run foreground commands later)
logs:
	docker compose logs -f dev
