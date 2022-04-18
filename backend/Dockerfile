FROM golang:1.16-alpine
WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .