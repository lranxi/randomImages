FROM golang:1.17.1-alpine

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main ./main.go

EXPOSE 8088

CMD /app/main