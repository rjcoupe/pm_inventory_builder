# syntax=docker/dockerfile:1

FROM golang:1.19.2-alpine3.16 AS builder

WORKDIR /app

COPY go.* .
RUN go mod download
COPY --link . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /out/pm_inventory_builder .

FROM alpine:3.16 AS app

COPY --link --chmod=755 --from=builder /out/pm_inventory_builder /
COPY --link --chmod=755 entrypoint.sh /

VOLUME /data

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "pm_inventory_builder", "-help" ]