FROM golang:tip-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./insights

EXPOSE 8080

CMD ["./insights"]
