FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/cypherium/cypherBFT && \
  (cd cypherBFT && make cypher) && \
  cp cypherBFT/build/bin/cypher /cypher && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /cypherBFT && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/cypher"]
