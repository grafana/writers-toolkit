FROM nginx:1.31.10-alpine-slim@sha256:c293b28729906d97fa5302fce3bd79bd27a3ea9a5d9604e21c6ff0ebf14549fc

RUN apk upgrade --no-cache

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY deploy-preview-files/deploy-preview/nginx.conf /etc/nginx/nginx.conf
COPY build.conf /etc/nginx/build.conf
COPY dist/redirects.conf /etc/nginx/redirects.conf

RUN nginx -t
