FROM alpine:latest as reviewdog

ENV REVIEWDOG_VERSION=0.14.1
ENV REVIEWDOG_CHECKSUM=fbd700763e119230e7c935ba27e0e65b9523c2d7c4f59c904670039a9010e3c86cb3778831eb483ff8e5816ac25dcde0bca7f0b0a155e408caa84d9e45249110

RUN mkdir /out

RUN wget -O reviewdog.tar.gz -q https://github.com/reviewdog/reviewdog/releases/download/v${REVIEWDOG_VERSION}/reviewdog_${REVIEWDOG_VERSION}_Linux_x86_64.tar.gz

RUN echo "${REVIEWDOG_CHECKSUM}  reviewdog.tar.gz" \
| sha512sum -csw - \
|| (printf "\nERROR: Invalid checksum\nexpected: %s\n     got: %s\n\n" "${REVIEWDOG_CHECKSUM}" "$(sha512sum reviewdog.tar.gz | awk '{print $1}')";  exit 1)

RUN tar x -zvf reviewdog.tar.gz reviewdog

RUN mv reviewdog /out/
RUN chmod +x /out/reviewdog

ENV TEMPLATE_CHECKSUM=904dffe57652c4610a48541c4fe9837953ea072d5fb1201ce06d74db0f82970fab4934a1089719c379512cfb955164a7e82eecca2398d33392be9802ff5a3bbf

RUN wget -O rdjsonl.tmpl -q https://raw.githubusercontent.com/errata-ai/vale-action/reviewdog/lib/rdjsonl.tmpl

RUN echo "${TEMPLATE_CHECKSUM}  rdjsonl.tmpl" \
| sha512sum -csw - \
|| (printf "\nERROR: Invalid checksum\nexpected: %s\n     got: %s\n\n" "${TEMPLATE_CHECKSUM}" "$(sha512sum rdjsonl.tmpl | awk '{print $1}')";  exit 1)

RUN mv rdjsonl.tmpl /out/

FROM jdkato/vale:v2.28.1

ENV DICPATH=/etc/vale/dictionaries

RUN mkdir -p /etc/vale

COPY Grafana /etc/vale/styles/Grafana
COPY .vale.ini /etc/vale
COPY dictionaries /etc/vale/dictionaries

# Git is useful for reporting with reviewdog.
RUN apk add --no-cache git

COPY --from=reviewdog /out/reviewdog /bin
COPY --from=reviewdog /out/rdjsonl.tmpl /etc/vale

RUN vale --config=/etc/vale/.vale.ini sync
