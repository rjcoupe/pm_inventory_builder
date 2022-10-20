# syntax=docker/dockerfile:1

FROM golang:1.19.2-alpine3.16 AS builder

WORKDIR /app

COPY --link . .
RUN go mod download
RUN go build -v -o pm_inventory_builder

FROM golang:1.19.2-alpine3.16 AS app

COPY --link --chmod=755 --from=builder /app/pm_inventory_builder /app/pm_inventory_builder
COPY --link --chmod=755 entrypoint.sh /entrypoint.sh

VOLUME /data

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "/app/pm_inventory_builder", "-help" ]