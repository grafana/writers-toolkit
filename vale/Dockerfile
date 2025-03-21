################################################################################
FROM --platform=${BUILDPLATFORM} golang:alpine AS inhibit-rules
ARG TARGETOS TARGETARCH

RUN mkdir /src /out
COPY /tools /src

WORKDIR /src
RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o inhibit-rules ./cmd/inhibit-rules
RUN mv inhibit-rules /out

################################################################################
FROM jdkato/vale:v3.8.0

COPY --from=inhibit-rules /out/inhibit-rules /bin

# Git is useful for diffing changed files.
RUN apk add --no-cache git

RUN mkdir -p /etc/vale
WORKDIR /etc/vale

COPY .vale.ini /etc/vale/.vale.ini
COPY Grafana /etc/vale/Grafana

RUN vale sync
