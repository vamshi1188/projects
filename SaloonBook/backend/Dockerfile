# Build stage
FROM golang:1.23-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd

# Production image
FROM alpine:3.20
WORKDIR /app
ENV PORT=5000
# Copy backend binary
COPY --from=build /app/server ./server
# Optionally copy frontend build if present (uncomment build step to include)
# COPY frontend/web/dist ./public
EXPOSE 5000
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD wget -qO- http://localhost:5000/api/health || exit 1
LABEL org.opencontainers.image.title="SaloonBook" \
      org.opencontainers.image.source="https://github.com/vamshi1188/SaloonBook" \
      org.opencontainers.image.description="Salon booking Go backend"
CMD ["./server"]
