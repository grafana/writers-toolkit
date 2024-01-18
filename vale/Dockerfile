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
RUN mv ./.reviewdog.yml /out

ENV TEMPLATE_CHECKSUM=904dffe57652c4610a48541c4fe9837953ea072d5fb1201ce06d74db0f82970fab4934a1089719c379512cfb955164a7e82eecca2398d33392be9802ff5a3bbf

RUN wget -O rdjsonl.tmpl -q https://raw.githubusercontent.com/errata-ai/vale-action/reviewdog/lib/rdjsonl.tmpl

RUN echo "${TEMPLATE_CHECKSUM}  rdjsonl.tmpl" \
| sha512sum -csw - \
|| (printf "\nERROR: Invalid checksum\nexpected: %s\n     got: %s\n\n" "${TEMPLATE_CHECKSUM}" "$(sha512sum rdjsonl.tmpl | awk '{print $1}')";  exit 1)

RUN mv rdjsonl.tmpl /out/

FROM --platform=$BUILDPLATFORM jdkato/vale:v3.0.5

ENV DICPATH=/etc/vale/dictionaries

RUN mkdir -p /etc/vale

COPY Google /etc/vale/styles/Google
COPY Grafana /etc/vale/styles/Grafana
COPY .vale.ini /etc/vale
COPY dictionaries /etc/vale/dictionaries

# Git is useful for reporting with reviewdog.
RUN apk add --no-cache git

COPY --from=reviewdog /out/reviewdog /bin
COPY --from=reviewdog /out/rdjsonl.tmpl /etc/vale
COPY --from=reviewdog /out/.reviewdog.yml /etc/vale

WORKDIR /etc/vale
RUN vale sync
