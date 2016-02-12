FROM scratch

MAINTAINER Ton Schoots <ton@maiastra.com>

EXPOSE 8080


COPY download_server .

ENTRYPOINT ["./download_server"]
