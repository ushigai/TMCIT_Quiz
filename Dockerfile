FROM golang:latest

WORKDIR /app
COPY . /app

#RUN go mod init main \
    #&& go mod tidy \
    #&& go build -o server

RUN go install
RUN go build -o server

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
EXPOSE 8080


CMD ["server"]

