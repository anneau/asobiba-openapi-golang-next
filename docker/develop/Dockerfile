FROM debian:latest

ENV GOLANG_VERSION 1.23.1

RUN apt-get update && apt-get install -y wget

RUN wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH $PATH:/usr/local/go/bin
