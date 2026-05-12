# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Build the binary (static, no CGO so we can run on a minimal base image)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# ---- Runtime stage ----
FROM alpine:3.20

WORKDIR /app

# Non-root user
RUN addgroup -S app && adduser -S app -G app

# Copy binary + runtime assets only (no Go toolchain, no sources)
COPY --from=builder /app/server ./server
COPY --chown=app:app data ./data
COPY --chown=app:app templates ./templates
COPY --chown=app:app static ./static

USER app

EXPOSE 8080

CMD ["./server"]
