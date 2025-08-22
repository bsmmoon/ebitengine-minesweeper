FROM golang:1.24-bookworm

# Core CLI tools only
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    curl \
    git \
  && rm -rf /var/lib/apt/lists/*

# Non-root user
RUN useradd -m dev
USER dev
WORKDIR /workspace

# Sensible Go env
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOTOOLCHAIN=auto \
    PATH="/home/dev/go/bin:${PATH}"
