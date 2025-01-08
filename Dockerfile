# Use official Golang image
FROM golang:1.22.5

# Set up working directory
WORKDIR /go/src/my-beego-app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy application files
COPY . .

# Install Beego CLI
RUN go install github.com/beego/bee/v2@latest

# Generate Swagger docs
RUN bee generate docs

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["bee", "run"]
