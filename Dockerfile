# ---------- Build stage ----------
FROM --platform=linux/amd64 golang:1.22 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely

# ---------- Runtime stage ----------
FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/notely /usr/bin/notely
CMD ["notely"]
