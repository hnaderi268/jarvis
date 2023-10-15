FROM golang:1.21
LABEL authors="hnaderi268"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /jarvis

EXPOSE 8080

ENTRYPOINT ["/jarvis"]