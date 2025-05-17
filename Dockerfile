FROM docker.io/library/golang:1.24.3-alpine3.21 AS builder
FROM scratch

# It ain't much, but it's honest work:
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import showmethatoken binary
COPY ./showmethatoken /usr/local/bin/showmethatoken

WORKDIR /root
ENTRYPOINT [ "/usr/local/bin/showmethatoken" ]
CMD []