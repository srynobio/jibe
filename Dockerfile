FROM ubuntu:latest
MAINTAINER Shawn Rynearson "shawn.rynearson@gmail.com"

LABEL vender="srynobio"
ENV VERSION 1.0.1 

RUN apt-get update && apt-get -y install \
	git \
	wget

RUN wget https://github.com/srynobio/jibe/releases/download/$VERSION/jibe_linux64
RUN mv jibe_linux64 jibe
RUN chmod 755 jibe
RUN mv jibe /usr/bin/

RUN jibe -version
