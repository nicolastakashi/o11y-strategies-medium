# Start from the official Golang image
FROM golang:1.20

# Create a directory inside the container to store the application code
RUN mkdir /app

# Set the working directory to the new directory
WORKDIR /app

# Copy the source code from the current directory to the new directory
COPY . .

RUN ls -la

# Compile the Go code into an executable binary
RUN go build -o main .

# Expose port 8081 for the web server to listen on
EXPOSE 8081

# Start the web server when the container starts
CMD ["/app/main"]