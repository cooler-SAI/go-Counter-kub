# Using official image Go
FROM golang:1.23-alpine

# Create working directory
WORKDIR /app

# Copy code to image
COPY . .

# Build the app
RUN go build -o go-Counter-kub .

# Open port 8081
EXPOSE 8081

# Run the app
CMD ["./go-Counter-kub"]