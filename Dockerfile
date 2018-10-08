FROM golang:1.11

WORKDIR /go/src/apiproject

COPY . .

RUN go get -d -v && go install -v

CMD ["sh", "-c", "/go/bin/apiproject"]

