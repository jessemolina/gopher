# ================================================================
# BUILD GO BINARY
FROM golang:1.19 as build_service-api
ENV CGO_ENABLED 0
ARG BUILD_REF

COPY . /service

WORKDIR /service/cmd/services/api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# ================================================================
# RUN BINARY IN ALPINE

FROM alpine:3.16
ARG BUILD_DATE
ARG BUILD_REF

RUN addgroup -g 1000 -S service && \
    adduser -u 1000 -h /service -G service -S service

COPY --from=build_service-api --chown=service:service /service/cmd/services/api/api /service/api

WORKDIR /service
USER service
CMD ["./api"]

# ================================================================
# LABEL

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="service" \
      org.opencontainers.image.authors="Jesse Molina"
