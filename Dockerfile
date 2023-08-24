# Use golang to build the app
FROM golang:1.21 AS build-stage
# Make the /app folder and cd into it
WORKDIR /app
# Get our go modules so that we can install them
COPY go.mod go.sum ./
# Download them
RUN go mod download && go mod tidy
# Move the go files here
ADD ./bin/SupergoonDiscordBot.tgz ./
# Build it
RUN CGO_ENABLED=0 GOOS=linux go build -o /discordBot

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
EXPOSE 80
COPY --from=build-stage /discordBot  /app
# EXPOSE 8000
ENTRYPOINT ["/app/discordBot"]