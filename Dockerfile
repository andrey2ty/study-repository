FROM golang:1.26.2-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go","run","main.go"]