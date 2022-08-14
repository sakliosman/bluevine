FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go build -v -o server


FROM alpine

COPY --from=builder /app/server /app/server

CMD ["/app/server"]
