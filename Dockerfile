FROM golang:1.5
MAINTAINER kchitta <kchitta@ebay.com>

ADD GoPlay /
RUN chmod +x /GoPlay
ENTRYPOINT ["/GoPlay"]
