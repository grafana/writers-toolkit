################################################################################
FROM --platform=${BUILDPLATFORM} golang:alpine@sha256:98e6cffc31ccc44c7c15d83df1d69891efee8115a5bb7ede2bf30a38af3e3c92 AS filter-sarif
ARG TARGETOS TARGETARCH

RUN mkdir /src /out
COPY /tools /src

WORKDIR /src
RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o filter-sarif ./cmd/filter-sarif
RUN mv filter-sarif /out

################################################################################
FROM jdkato/vale:v3.13.1@sha256:5c8e1a491e9523d71b864ce7e2ed71c5cdb96a4fb7bf7ae54c42e0786c8dba99

COPY --from=filter-sarif /out/filter-sarif /bin

# Git is useful for diffing changed files.
RUN apk add --no-cache git

RUN mkdir -p /etc/vale
WORKDIR /etc/vale

COPY .vale.ini /etc/vale/.vale.ini
COPY sarif.tmpl /etc/vale/sarif.tmpl
COPY Grafana /etc/vale/Grafana

RUN vale sync
