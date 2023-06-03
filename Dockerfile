FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 5555

RUN go build -o main .

CMD ["/app/main"]
