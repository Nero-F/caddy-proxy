FROM jwilder/docker-gen:debian AS docker-gen

FROM caddy:2.10.0-alpine

RUN apk add --no-cache --virtual .run-deps bash curl

WORKDIR /app

COPY --from=docker-gen /usr/local/bin/docker-gen /usr/local/bin/docker-gen
COPY docker-entrypoint.sh  entrypoint.sh
COPY caddy.tmpl docker-gen.cfg .

ENV DOCKER_HOST=unix:///tmp/docker.sock

EXPOSE 80
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
