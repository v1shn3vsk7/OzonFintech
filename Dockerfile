FROM golang:latest
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -a /app/cmd/server/
ENTRYPOINT exec go run cmd/server/main.go