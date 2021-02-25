FROM golang:1.14.0-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
EXPOSE 1323
CMD ["/app/main"]