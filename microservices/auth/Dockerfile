FROM golang:1.19

# Port to expose
ENV PORT=8080

WORKDIR /go/src/app

# Install libvips
RUN apt-get update && apt-get install -y libvips-dev

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Install go dependencies
RUN go mod download

# Set PKG_CONFIG_PATH for libvips
ENV PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
RUN pkg-config --libs vips

# Copy the source code
COPY . .

# Uncomment to build the app
# Build the app
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./main"]