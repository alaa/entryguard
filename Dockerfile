FROM alpine:latest

COPY entryguard /usr/local/bin/
ENTRYPOINT /usr/local/bin/entryguard
