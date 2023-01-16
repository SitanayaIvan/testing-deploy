FROM golang:1.19-alpine3.16 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/backend .

FROM alpine:3.16
COPY --chown=65534:65534 --from=builder /go/bin/backend .
ENV PORT=8080

EXPOSE 8080
CMD [ "./backend" ]