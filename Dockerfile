FROM alpine:3.15.0

COPY tvctl /usr/local/bin/tvctl
RUN chmod +x /usr/local/bin/tvctl

RUN mkdir /workdir
WORKDIR /workdir

ENTRYPOINT [ "/usr/local/bin/tvctl" ]
