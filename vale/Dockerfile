################################################################################
FROM --platform=${BUILDPLATFORM} golang:alpine@sha256:26111811bc967321e7b6f852e914d14bede324cd1accb7f81811929a6a57fea9 AS filter-sarif
ARG TARGETOS TARGETARCH

RUN mkdir /src /out
COPY /tools /src

WORKDIR /src
RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o filter-sarif ./cmd/filter-sarif
RUN mv filter-sarif /out

################################################################################
FROM jdkato/vale:v3.13.0@sha256:04bb8794dff4505bcf588522e860ba9547b73f02dd8b2de0faecb155472da672

COPY --from=filter-sarif /out/filter-sarif /bin

# Git is useful for diffing changed files.
RUN apk add --no-cache git

RUN mkdir -p /etc/vale
WORKDIR /etc/vale

COPY .vale.ini /etc/vale/.vale.ini
COPY sarif.tmpl /etc/vale/sarif.tmpl
COPY Grafana /etc/vale/Grafana

RUN vale sync
