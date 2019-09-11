FROM golang:1.13.0-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV PORT=8080

WORKDIR /go/src/github.com/Tech4GoodPH/go-sample-api
COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build -o ./build/go-sample-api -v .

EXPOSE ${PORT}

CMD [ "./build/go-sample-api" ]