#  Base Image
FROM golang:1.23.3-alpine AS builder 
# distroless
# builder is used to compile the go language
WORKDIR /app
# Copy go mod and sum files for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . . 

# Build the Go application
RUN go build -o productsapi ./cmd/productsapi

# # Stage 2: Create a minimal final image
FROM alpine:latest AS runner 

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/productsapi /app/

# ENV MONGODB_URI=mongodb://localhost:27017
# ENV POSTGRES_URI=postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable

EXPOSE 8080

CMD ["./productsapi"]
# ENTRYPOINT [ "GO", "cmd/productsapi/main.go"]