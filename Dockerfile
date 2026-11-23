FROM golang:1.27.4

LABEL maintainer="Tselmuun"

WORKDIR /server

COPY .. ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o server .

EXPOSE 8000

CMD ["./server"]
