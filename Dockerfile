FROM golang:latest

WORKDIR D:/GIT/Golang
#RUN ls

# Fetch dependencies
COPY go.mod ./
RUN go mod download

# Build
COPY ./ ./
RUN go build -o auth cmd/main.go


# Create final image
CMD ["./auth"]