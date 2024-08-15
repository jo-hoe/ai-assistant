FROM golang:1.23-bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates \
    gcc \
    make \
    libgtk-3-dev libwebkit2gtk-4.0-dev \
    && rm -rf /var/lib/apt/lists/*
    
RUN wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.58.1

WORKDIR /app
COPY ../. /app
