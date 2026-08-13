FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bot .

EXPOSE 80

ENTRYPOINT ["./bot"]
