FROM golang:1.16.2-alpine3.13
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o ./output/bin/main ./cmd/app/
CMD ["/app/output/bin/main"]
