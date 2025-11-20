# ---------- Build Stage ----------
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ai-summary .

# ---------- Run Stage ----------
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/ai-summary .

ENV PORT=8080

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

CMD ["./ai-summary"]
