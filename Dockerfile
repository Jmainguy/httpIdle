# Dockerfile
FROM alpine:latest
COPY httpIdle /usr/bin/httpIdle
EXPOSE 8080
CMD ["/usr/bin/httpIdle"]
