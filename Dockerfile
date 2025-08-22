# FROM golang:1.22-bookworm
FROM golang:1.24-bookworm

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    git \
    xorg-dev libgl1-mesa-dev libglu1-mesa-dev libasound2-dev \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m dev
USER dev
WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOTOOLCHAIN=auto \
    PATH="/home/dev/go/bin:${PATH}"
