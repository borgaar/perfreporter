FROM golang:alpine3.21 AS builder

WORKDIR /build

RUN addgroup --system --gid 1001 builder && \
    adduser --system --uid 1001 builder && \
    chown -R 1001:1001 /build

USER builder

COPY --chown=1001:1001 . .

RUN go build

FROM golang:alpine3.21 AS runner

WORKDIR /server

RUN addgroup --system --gid 1001 runner && \
    adduser --system --uid 1001 runner && \
    chown -R 1001:1001 /server

USER runner

COPY --from=builder --chown=1001:1001 /build/perfreporter perfreporter
ENV PORT=8080

EXPOSE 8080

CMD ["./perfreporter"]
