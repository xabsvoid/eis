FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

ADD https://github.com/xabsvoid/eis.git .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o eis ./cmd/eis

FROM alpine:3.22.0

COPY --from=build /app/eis /eis

ENTRYPOINT ["./eis", "-dsn", "${EIS_HOST}"]
