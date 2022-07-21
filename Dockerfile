FROM golang:1.18 as builder
ARG VERSION
ARG SHORT_COMMIT
ARG DATE
COPY . /tvctl
WORKDIR /tvctl
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=$VERSION -X main.commit=$SHORT_COMMIT -X main.date=$DATE" -o tvctl ./cmd/tvctl/main.go

FROM golang:1.18
COPY --from=builder /tvctl/tvctl /usr/bin/
CMD ["tvctl"]
