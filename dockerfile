FROM golang:1.19

WORKDIR /SE_MIM22_WEBSHOP_ORDERSERVICE

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8443

RUN go mod download

ENTRYPOINT go build  && ./SE_MIM22_WEBSHOP_ORDERSERVICE