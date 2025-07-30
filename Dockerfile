FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080
CMD ["./main"]
# Use the following line if you want to run the application directly
# CMD ["go", "run", "./cmd/main.go"]
