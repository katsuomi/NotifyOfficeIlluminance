# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir -p $GOPATH/src/github.com/katsuomi/NotifyOfficeIlluminance
# コンテナログイン時のディレクトリ指定
WORKDIR $GOPATH/src/github.com/katsuomi/NotifyOfficeIlluminance
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . $GOPATH/src/github.com/katsuomi/NotifyOfficeIlluminance
# 必要なパッケージをイメージにインストールする
RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres
