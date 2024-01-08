ARG BASE_IMAGE=ishocon2-app-base:latest
FROM ${BASE_IMAGE}

RUN sudo apt-get update && \
    sudo apt-get install -y percona-toolkit && \
    sudo apt-get clean

# Go のインストール
ARG TARGETARCH
RUN sudo wget -q https://dl.google.com/go/go1.21.5.linux-${TARGETARCH}.tar.gz && \
    sudo tar -C /usr/local -xzf go1.21.5.linux-${TARGETARCH}.tar.gz && \
    sudo rm go1.21.5.linux-${TARGETARCH}.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOROOT /usr/local/go
ENV GOPATH /home/ishocon/.local/go
ENV PATH $PATH:$GOROOT/bin

# アプリケーション
COPY --chown=ishocon:ishocon webapp/ /home/ishocon/webapp
RUN cd ~/webapp/go && \
    mkdir -p /tmp/go && \
    GOARCH=${TARGETARCH} go build -x -o /tmp/go/webapp *.go

WORKDIR /home/ishocon

# スクリプト配置
COPY --chown=ishocon:ishocon --chmod=777 scripts/ /home/ishocon/scripts/

# 頻繁に書き換えるファイルのコピー
COPY admin/config/my.app.cnf /etc/mysql/my.cnf
COPY admin/config/nginx.conf /etc/nginx/nginx.conf

COPY run.sh /home/ishocon/run.sh
