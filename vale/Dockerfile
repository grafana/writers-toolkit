################################################################################
FROM --platform=${BUILDPLATFORM} golang:alpine@sha256:4c9fe60190a2a3350ddc51de80d0224b8a6698d12bdfc999fee45ea9d6c46dbc AS filter-sarif
ARG TARGETOS TARGETARCH

RUN mkdir /src /out
COPY /tools /src

WORKDIR /src
RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o filter-sarif ./cmd/filter-sarif
RUN mv filter-sarif /out

################################################################################
FROM jdkato/vale:v3.21.0@sha256:8a5a05b5c1751c6a6132ca459d67ffe3ad2cc338abc5795ee634470dd39de9c4

COPY --from=filter-sarif /out/filter-sarif /bin

# Git is useful for diffing changed files.
RUN apk add --no-cache git

RUN mkdir -p /etc/vale
WORKDIR /etc/vale

COPY .vale.ini /etc/vale/.vale.ini
COPY sarif.tmpl /etc/vale/sarif.tmpl
COPY Grafana /etc/vale/Grafana

RUN vale sync
