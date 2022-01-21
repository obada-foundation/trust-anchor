FROM alpine:edge AS builder

RUN apk update && apk upgrade

RUN apk add --no-cache go

COPY backend/ /src

RUN cd /src && \
    CGO_ENABLED=1 GOOS=linux go build -ldflags '-extldflags "-static"' -a -o trust-anchor main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates && \
    adduser -D ta && \
    mkdir -p /var/obada/trust-anchor/data && \
    chown -R ta:ta /var/obada/trust-anchor

USER ta

WORKDIR /home/ta

COPY --from=builder --chown=ta /src/trust-anchor .

CMD ["/home/ta/trust-anchor", "server"]
