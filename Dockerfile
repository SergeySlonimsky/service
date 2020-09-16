ARG golangVer=1.14
ARG alpineVer=3.11

FROM golang:${golangVer}-alpine${alpineVer} as builder

WORKDIR /var/app
COPY . ./
RUN go mod download
RUN go build -v -o app

RUN chmod +x app

FROM alpine:${alpineVer}

WORKDIR /var/app/
COPY --from=builder /var/app/app .
ENTRYPOINT ["/var/app/app"]
