FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

RUN CGO_ENABLED=0 go build

FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /app

COPY --from=builder /app/trainspotter-backend .

EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["./trainspotter-backend"]
CMD []