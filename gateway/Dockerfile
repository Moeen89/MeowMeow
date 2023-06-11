# Use a base Go image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod go.sum ./

# Download and cache Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Create a new stage for the minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/app .

# Expose the gRPC port (replace with the appropriate port number)
EXPOSE 6433

# Run the Go application
CMD ["./app"]
