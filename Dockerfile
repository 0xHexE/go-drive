FROM golang:latest as builder
LABEL maintainer="Omkar Yadav <httpsOmkar@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Built the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

# Install the SSL certificates and curl
RUN apk --no-cache add ca-certificates curl

WORKDIR /root/
COPY --from=builder /app/app .
RUN find . -name "*.go" -type f -delete

# Health check for the app
HEALTHCHECK --interval=10s --timeout=3s \
CMD curl --fail http://localhost:$PORT/_health || exit 1

EXPOSE $PORT
USER 1000
CMD ["./app"]
