FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

ADD https://github.com/pressly/goose.git#v3.25.0 .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o goose ./cmd/goose

FROM alpine:3.22.0

COPY --from=build /app/goose /goose

ENTRYPOINT ["/goose", "up"]
