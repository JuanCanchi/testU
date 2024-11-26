FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/cmd/users/createUsers cmd/users/users.go
RUN go build -o /app/cmd/api/main cmd/api/main.go

FROM golang:1.20

WORKDIR /app

COPY --from=builder /app/cmd/users/createUsers /app/cmd/users/createUsers
COPY --from=builder /app/cmd/api/main /app/cmd/api/main

EXPOSE 8080

CMD ["/app/cmd/api/main"]
