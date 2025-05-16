# ---------- Builder Stage ----------
FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o insights

# ---------- Final Image ----------
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/insights .

EXPOSE 8080

CMD ["./insights"]