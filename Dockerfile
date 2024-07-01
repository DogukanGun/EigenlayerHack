FROM ubuntu:latest

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

WORKDIR /image
ADD main /image

ENTRYPOINT ./main

