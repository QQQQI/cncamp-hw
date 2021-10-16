FROM golang:alpine3.14
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o httpserver httpserver.go
EXPOSE 8000
ENTRYPOINT ./httpserver
