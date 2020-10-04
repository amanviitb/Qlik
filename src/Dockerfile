FROM golang:1.15-alpine

ENV BINARY_NAME="palindrome-message"

WORKDIR /go/src/app

COPY ./go.mod ./
COPY ./go.sum ./

# download all the dependencies
RUN go mod download

# copy all the files into the container image filesnapshot
COPY . .

RUN go build -o ../bin/${BINARY_NAME} .

CMD ["../bin/palindrome-message"]


