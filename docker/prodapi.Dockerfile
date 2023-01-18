# stage 1: building application binary file
FROM golang:1.19-alpine as build

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -mod vendor -o main main.go

# stage 2: copy only the application binary file and necessary files to the alpine container
FROM alpine:latest
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=build /app/main .

# run the service on container startup.
CMD ["/app/main"]