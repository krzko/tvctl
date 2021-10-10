FROM golang:1.17.2-alpine3.13 AS builder
WORKDIR /src/
COPY . /src/
ARG COMMIT
ARG NOW
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/tvctl -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/tvctl/main.go

FROM scratch
ARG COMMIT
ARG NOW
ARG VERSION
LABEL maintainer="Kristof Kowalski <k@ko.wal.ski>" \
    org.opencontainers.image.title="tvctl" \
    org.opencontainers.image.description="A command-line utility to interact with TradingView" \
    org.opencontainers.image.authors="Kristof Kowalski <k@ko.wal.ski>" \
    org.opencontainers.image.vendor="Kristof Kowalski" \
    org.opencontainers.image.documentation="https://github.com/krzko/tvctl/docs" \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.version=$VERSION \
    org.opencontainers.image.url="https://ko.wal.ski" \
    org.opencontainers.image.source="https://github.com/krzko/tvctl.git" \
    org.opencontainers.image.revision=$COMMIT \
    org.opencontainers.image.created=$NOW
COPY --from=builder /bin/tvctl /bin/tvctl
ENTRYPOINT ["/bin/tvctl"]
