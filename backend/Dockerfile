FROM golang:1.20.6-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY handlers ./handlers
COPY models ./models
COPY db ./db
COPY utils ./utils

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

EXPOSE 8080

# Run
ENTRYPOINT ["./url-shortener-backend"]