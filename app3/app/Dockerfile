## 2020/10/14最新versionを取得
#FROM golang:1.15.2-alpine
## アップデートとgitのインストール
#RUN apk update && apk add git
## appディレクトリの作成
#RUN mkdir /go/src/app
## ワーキングディレクトリの設定
#WORKDIR /go/src/app
## ホストのファイルをコンテナの作業ディレクトリに移行
#ADD . /go/src/app


#RUN go get io\
    #go get -u github.com/oxequa/realize \
    #&& go get github.com/go-sql-driver/mysql

#CMD ["realize", "start"]
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


CMD ["/app/server"]
