# Step 1: Use official Golang image as the base
FROM golang:1.20-alpine AS build

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Step 4: Download the Go module dependencies
RUN go mod download

# Step 5: Copy the rest of the application code to the working directory
COPY . .

# Step 6: Build the application
RUN go build -o metafetch .

# Step 7: Use a smaller base image for the final stage
FROM alpine:latest

# Step 8: Set working directory in the final stage
WORKDIR /app

# Step 9: Copy the built binary from the builder stage
COPY --from=build /app/metafetch .

# Step 10: Expose the application port (adjust if necessary)
EXPOSE 8080

# Step 11: Command to run the application
CMD ["./metafetch"]
