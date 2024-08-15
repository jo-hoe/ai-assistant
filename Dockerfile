FROM golang:1.23-bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates \
    gcc \
    make \
    libgtk-3-dev libwebkit2gtk-4.0-dev \
    && rm -rf /var/lib/apt/lists/*
    
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1

WORKDIR /app
COPY ../. /app
