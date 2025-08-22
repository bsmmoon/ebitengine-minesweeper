#!/usr/bin/env bash
set -euo pipefail

# Show env for sanity
go env

# Pin versions compatible with Go 1.24+
go install golang.org/x/tools/gopls@v0.20.0
go install honnef.co/go/tools/cmd/staticcheck@2024.1.1
go install github.com/go-delve/delve/cmd/dlv@v1.22.0

echo "✅ Installed: gopls v0.20.0, staticcheck 2024.1.1, dlv v1.22.0"
