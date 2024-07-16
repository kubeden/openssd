# Build stage
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy templates and static files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Environment variables
ENV GITHUB_USERNAME=kubeden
ENV GITHUB_REPO=kubeden
ENV README_FILE=README.md
ENV INFO_FILE=INFO.md
ENV USER_NAME=Kuberdenis
ENV TEMPLATE_CHOICE=ssi

# Command to run the executable
CMD ["./main"]