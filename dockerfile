FROM golang:1.19-alpine

WORKDIR /SE_MIM22_WEBSHOP_ORDERSERVICE

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8460

RUN go mod download

ENTRYPOINT go build  && ./SE_MIM22_WEBSHOP_ORDERSERVICE