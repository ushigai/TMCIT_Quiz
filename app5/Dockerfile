FROM golang:latest

WORKDIR /app
COPY . /app

#RUN go mod tidy
#RUN go get github.com/cosmtrek/air
#CMD ["air"]

RUN go install && go build -o server

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
EXPOSE 1323


CMD ["/app/server"]
#CMD ["go", "run", "server.go"]

