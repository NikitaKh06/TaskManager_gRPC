FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o task-manager-api-service-grpc .

CMD ["./task-manager-api-service-grpc"]