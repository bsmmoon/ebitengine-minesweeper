#!/usr/bin/env bash
set -euo pipefail

# Ensure module cache is sane
go env

# Core dev tools
go install golang.org/x/tools/gopls@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/go-delve/delve/cmd/dlv@latest

# Optional: tidy modules if a go.mod exists
if [ -f go.mod ]; then
  go mod tidy
fi

echo "✅ Dev tools installed (gopls, staticcheck, dlv)"
