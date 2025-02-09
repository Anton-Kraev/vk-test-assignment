FROM golang:1.23 AS build

WORKDIR /cmd

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app -v ./cmd/app

FROM alpine:3.20 AS final

WORKDIR /

COPY --from=build /bin/app /app

EXPOSE 8080
