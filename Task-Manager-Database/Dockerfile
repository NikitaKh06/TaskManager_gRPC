FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o task-manager-database-service-grpc .

CMD ["./task-manager-database-service-grpc"]