FROM --platform=$BUILDPLATFORM golang:alpine as reviewdog
ARG TARGETOS TARGETARCH

ENV REVIEWDOG_VERSION=0.16.0
ENV REVIEWDOG_CHECKSUM=a847b248b045b9706cffb0cb34633476f86a79ab663db3ebfce3707105f7911ded5fccb5d0f632708c698cf3b9b361705de7c9ebbceecf71ad4503f2b787092d

RUN mkdir /src /out

RUN wget -O reviewdog.tar.gz -q "https://github.com/reviewdog/reviewdog/archive/refs/tags/v${REVIEWDOG_VERSION}.tar.gz"

RUN echo "${REVIEWDOG_CHECKSUM}  reviewdog.tar.gz" \
| sha512sum -csw - \
|| (printf "\nERROR: Invalid checksum\nexpected: %s\n     got: %s\n\n" "${REVIEWDOG_CHECKSUM}" "$(sha512sum reviewdog.tar.gz | awk '{print $1}')";  exit 1)

RUN tar zxf reviewdog.tar.gz -C /src

WORKDIR "/src/reviewdog-${REVIEWDOG_VERSION}"
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o reviewdog ./cmd/reviewdog
RUN mv reviewdog /out

FROM --platform=${TARGETPLATFORM} jdkato/vale:v3.0.5

ENV DICPATH=/etc/vale/dictionaries

RUN mkdir -p /etc/vale

COPY Google /etc/vale/styles/Google
COPY Grafana /etc/vale/styles/Grafana
COPY .vale.ini /etc/vale/.vale.ini
COPY .vale-config /etc/vale/.vale-config
COPY dictionaries /etc/vale/dictionaries
COPY rdjsonl.tmpl /etc/vale/rdjsonl.tmpl
COPY .reviewdog.yml /etc/vale/.reviewdog.yml

# Git is useful for reporting with reviewdog.
RUN apk add --no-cache git

COPY --from=reviewdog /out/reviewdog /bin

WORKDIR /etc/vale
RUN vale sync
