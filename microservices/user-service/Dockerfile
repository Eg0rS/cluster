FROM golang:1.19-alpine

EXPOSE 80

WORKDIR /go/src/user-service
COPY . .
RUN chmod +x ./start.sh

RUN go install -mod vendor

ENTRYPOINT ./start.sh