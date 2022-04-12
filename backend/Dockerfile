FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# ENV DB_ADDR="root:admin@tcp(172.17.0.1:3306)/igeargeekpracticdb?charset=utf8mb4&parseTime=True&loc=Local" JWT_SECRET_KEY="igeargeekjiw" PORT="8000"