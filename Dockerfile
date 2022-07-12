FROM golang:alpine AS builder

LABEL go.package.version="0.1" \
    go.package.release.date="2022-07-12"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod main.go ./
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main .
EXPOSE 9999
ENTRYPOINT ["/main"]

# multi stage

## build
# docker build . -t go-docker
# docker run -d -p 9999:9999 go-docker