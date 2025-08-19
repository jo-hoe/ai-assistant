FROM golang:1.25-bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates \
    gcc \
    make \
    libgtk-3-dev libwebkit2gtk-4.0-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . /app
