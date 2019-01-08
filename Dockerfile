FROM ubuntu:latest
MAINTAINER Shawn Rynearson "shawn.rynearson@gmail.com"

LABEL vender="srynobio"

RUN apt-get update && apt-get -y install \
	git \
	wget


RUN wget https://github.com/srynobio/jibe/releases/download/1.0.0/jibe_linux64
RUN mv jibe_linux64 jibe
RUN chmod 755 jibe
RUN mv jibe /usr/bin/

RUN jibe -version
