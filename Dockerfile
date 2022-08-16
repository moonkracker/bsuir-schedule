FROM golang:1.19-alpine3.16 AS builder
WORKDIR /opt/build
ARG VERSION_TAG
COPY . .
RUN apk add build-base && go mod tidy && \
    go build -o ./bsuir-schedule -ldflags "-X 'main.Version=${VERSION_TAG}'"
HEALTHCHECK --interval=5s --timeout=10s --retries=3 CMD ./bsuir-schedule --version || exit 1

FROM alpine:3.16
LABEL maintainer="Maksim Kananovich"
COPY --from=builder /opt/build/bsuir-schedule /usr/bin
ENTRYPOINT ["bsuir-schedule"]