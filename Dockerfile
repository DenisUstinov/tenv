# Modules Caching
FROM golang:1.23 AS modules

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Build
FROM golang:1.23

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /my-app ./examples

ENTRYPOINT ["/my-app"]