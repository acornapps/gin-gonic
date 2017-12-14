FROM golang:1.9.2 
RUN mkdir -p /go/src/app 
ADD . /go/src/app

WORKDIR /go/src/app 

RUN go build -o main .

EXPOSE 8080

CMD ["/go/src/app/main"]

