FROM golang:1.12.1-stretch

WORKDIR /src
COPY . /src

RUN go build -o go-note-you

EXPOSE 1323
CMD ["./go-note-you"]