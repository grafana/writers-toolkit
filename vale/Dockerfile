################################################################################
FROM --platform=${BUILDPLATFORM} golang:alpine@sha256:d200d2dd76b03a99b52651fd1325aff7dcd35422c35c01bac56571b0507f3deb AS filter-sarif
ARG TARGETOS TARGETARCH

RUN mkdir /src /out
COPY /tools /src

WORKDIR /src
RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o filter-sarif ./cmd/filter-sarif
RUN mv filter-sarif /out

################################################################################
FROM jdkato/vale:v3.26.0@sha256:d4ebc402b255dacc759c468ed8c91bd5be82918bbb59583558539f3be80f1f01

COPY --from=filter-sarif /out/filter-sarif /bin

# Git is useful for diffing changed files.
RUN apk add --no-cache git

RUN mkdir -p /etc/vale
WORKDIR /etc/vale

COPY .vale.ini /etc/vale/.vale.ini
COPY sarif.tmpl /etc/vale/sarif.tmpl
COPY Grafana /etc/vale/Grafana

RUN vale sync
