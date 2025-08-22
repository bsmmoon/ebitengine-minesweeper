SHELL := /bin/bash

.PHONY: build up sh web wasm clean

# Build the dev image
build:
	docker compose build

# Start the container (idle)
up:
	docker compose up -d

# Shell into the container
sh:
	docker compose exec dev bash

# Build WASM and run a tiny server on :8080
web: wasm
	docker compose exec -d dev bash -lc 'go run ./tools/serve'

# Compile to WASM (outputs to ./web/)
wasm:
	docker compose exec dev bash -lc '\
	  mkdir -p web && \
	  cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/ && \
	  GOOS=js GOARCH=wasm go build -o web/game.wasm ./cmd/game \
	'

# Cleanup generated artifacts
clean:
	rm -rf web
